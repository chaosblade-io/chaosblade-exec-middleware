/*
 * Copyright 1999-2020 Alibaba Group Holding Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package nginx

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

type testFuncType func(context.Context, *testCase)

const testConfig = "nginx.conf"
const cleanTimeout = time.Duration(1) * time.Second
const mockUid = "mock"

type testCase struct {
	spec        spec.ExpActionCommandSpec
	t           *testing.T
	channel     spec.Channel
	client      http.Client
	mustHaveLua bool
	testFunc    testFuncType
}

func newTestCase(t *testing.T, testFunc testFuncType, spec spec.ExpActionCommandSpec, mustHaveLua bool) *testCase {
	tc := &testCase{
		spec:        spec,
		t:           t,
		testFunc:    testFunc,
		channel:     channel.NewLocalChannel(),
		client:      http.Client{},
		mustHaveLua: mustHaveLua,
	}
	tc.spec.Executor().SetChannel(tc.channel)
	return tc
}

func (tc *testCase) prepare(ctx context.Context) error {
	resp := tc.channel.Run(ctx, "nginx -v", "")
	if !resp.Success ||
		(tc.mustHaveLua && !strings.Contains(resp.Result.(string), "openresty")) {
		return errors.New("nginx-openresty not exists, test case prepare failed")
	}
	if resp = testNginxExists(tc.channel, ctx); resp != nil {
		return errors.New("nginx not exists, test case prepare failed")
	}
	if resp := swapNginxConfig(tc.channel, ctx, testConfig, nil); resp != nil && !resp.Success {
		return errors.New(resp.Err)
	}
	return nil
}

func (tc *testCase) clean() {
	ctx, cancel := context.WithTimeout(context.Background(), cleanTimeout)
	defer cancel()
	tc.destroy(ctx)
	reloadNginxConfig(tc.channel, ctx)
}

func (tc *testCase) expectResponse(ctx context.Context, path, respBody string, respCode, port int, respHeaders map[string]string) {
	// 'nginx -s reload' is asynchronous, so an immediate http request may return the old response.
	killNginx(tc.channel, ctx)
	startNginx(tc.channel, ctx)
	resp, err := tc.client.Get(fmt.Sprintf("http://localhost:%v%s", port, path))
	if err != nil {
		tc.t.Fatal(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	bodyString := strings.TrimSpace(string(body[:]))
	if err != nil {
		tc.t.Fatal(err.Error())
	}
	if resp.StatusCode != respCode {
		tc.t.Fatalf("response code not match, expect %v but actual %v", respCode, resp.StatusCode)
	}
	if respBody != "" && bodyString != respBody {
		tc.t.Fatalf("response body not match, expect %v but actual %v", respBody, bodyString)
	}
	for k, v := range respHeaders {
		match := false
		if vv, ok := resp.Header[k]; ok {
			for _, i := range vv {
				if v == i {
					match = true
					break
				}
			}
		}
		if !match {
			tc.t.Fatalf("response headers not match: %v=%v not match", k, v)
		}
	}
}

func (tc *testCase) expectAlive(ctx context.Context, isAlive bool) {
	alive := true
	if resp := testNginxExists(tc.channel, ctx); resp != nil {
		alive = false
	}
	if isAlive != alive {
		tc.t.Fatal("alive state not match")
	}
}

func (tc *testCase) start(ctx context.Context, args map[string]string) {
	model := spec.ExpModel{}
	model.ActionFlags = args
	if resp := tc.spec.Executor().Exec(mockUid, ctx, &model); !resp.Success {
		tc.t.Fatalf("cmd '%v' run failed:%s", args, resp.Err)
	}
}

func (tc *testCase) destroy(ctx context.Context) {
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	tc.spec.Executor().Exec(mockUid, context.WithValue(ctx, "suid", mockUid), &model)
}

func (tc *testCase) test(ctx context.Context) {
	defer tc.clean()
	if err := tc.prepare(ctx); err != nil {
		tc.t.Skipf("Skip test: %s", err.Error())
	}
	tc.testFunc(ctx, tc)
}

func testWithTimeout(t *testing.T, testFunc testFuncType, spec spec.ExpActionCommandSpec, mustHaveLua bool, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	newTestCase(t, testFunc, spec, mustHaveLua).test(ctx)
}

func TestNginxCrash(t *testing.T) {
	testWithTimeout(t, func(ctx context.Context, tc *testCase) {
		tc.expectAlive(ctx, true)
		tc.start(ctx, map[string]string{})
		tc.expectAlive(ctx, false)
		tc.destroy(ctx)
		tc.expectAlive(ctx, true)
		tc.start(ctx, map[string]string{})
	}, NewCrashActionSpec(), false, time.Duration(1)*time.Second)
}

func TestNginxConfigChange(t *testing.T) {
	testWithTimeout(t, func(ctx context.Context, tc *testCase) {
		tc.expectResponse(ctx, "/", "", 200, 80, nil)

		tc.start(ctx, map[string]string{"mode": "cmd", "set-config": "listen=9999", "block": "http.server[0]"})
		tc.expectResponse(ctx, "/", "", 200, 9999, nil)

		tc.destroy(ctx)
		tc.expectResponse(ctx, "/", "", 200, 80, nil)
	}, NewConfigActionSpec(), false, time.Duration(2)*time.Second)
}

func TestNginxResponseChange(t *testing.T) {
	testWithTimeout(t, func(ctx context.Context, tc *testCase) {
		tc.expectResponse(ctx, "/test", "", 404, 80, nil)
		tc.expectResponse(ctx, "/", "", 200, 80, nil)

		tc.start(ctx, map[string]string{"regex": "/t.*", "code": "200", "body": "ok", "header": "Server=mock;"})
		tc.expectResponse(ctx, "/test", "ok", 200, 80, map[string]string{"Server": "mock"})
		tc.expectResponse(ctx, "/tt", "ok", 200, 80, map[string]string{"Server": "mock"})

		tc.start(ctx, map[string]string{"path": "/path", "code": "200", "body": "path"})
		tc.expectResponse(ctx, "/path", "path", 200, 80, nil)

		tc.destroy(ctx)
		tc.expectResponse(ctx, "/test", "", 404, 80, nil)
	}, NewResponseActionSpec(), true, time.Duration(2)*time.Second)
}
