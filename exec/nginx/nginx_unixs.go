//go:build !windows
// +build !windows

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
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

// Copy nginx.conf to nginx config path and test it.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string, nginxPath string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf("cp %s %s", file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, nginxPath, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("rm %s", tmpFile), "") // ignore response
	if !response.Success || !strings.Contains(response.Result.(string), "successful") {
		return response
	}
	return nil
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	response := channel.Run(ctx,
		`ps aux | grep -v grep | egrep 'nginx: ' | awk '{print $2}'`, "")
	if !response.Success {
		return response
	}
	result := response.Result.(string)
	if strings.Count(result, "\n") == 0 {
		return spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx process")
	}
	return nil
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	if resp := testNginxExists(channel, ctx); resp != nil {
		return resp
	}

	if resp := channel.Run(ctx, "killall -9 nginx", ""); !resp.Success {
		return resp
	}
	return nil
}

func runNginxCommand(channel spec.Channel, ctx context.Context, nginxPath string, args string) *spec.Response {
	// "/usr/local/nginx/sbin/nginx"
	return channel.Run(ctx, nginxPath, args)
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return channel.Run(ctx, fmt.Sprintf("mv -f %s %s", backup, activeFile), "")
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	cmd := ""
	if util.IsExist(backup) {
		// don't create new backup
		cmd = fmt.Sprintf("cp -f %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("cp %s %s && cp -f %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && rm %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}
