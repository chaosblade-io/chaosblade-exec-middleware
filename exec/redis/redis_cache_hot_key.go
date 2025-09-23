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

package redis

import (
	"bytes"
	context "context"
	"fmt"
	"regexp"
	"strconv"
	"sync"
	time "time"

	"github.com/go-redis/redis/v8"
	"github.com/shirou/gopsutil/v3/process"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const (
	CacheHotKeyBin         = "chaos_cacheHotKey"
	DefaultThreadCount     = "20"
	DefaultSize            = "512"
	DefaultDuration        = "5m"
	ActionCommandName      = "cache-hot-key"
	ShortActionCommandName = "chk"
)

type CacheHotKeyActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCacheHotKeyActionSpec() spec.ExpActionCommandSpec {
	return &CacheHotKeyActionCommandSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "addr",
					Desc: "The address of redis server",
				},
				&spec.ExpFlag{
					Name: "password",
					Desc: "The password of redis server",
				},
				&spec.ExpFlag{
					Name: "key",
					Desc: "The key to be used",
				},
				&spec.ExpFlag{
					Name:    "thread-count",
					Desc:    "thread-count to do get and set command, default is 20",
					Default: DefaultThreadCount,
				},
				&spec.ExpFlag{
					Name:    "size",
					Desc:    fmt.Sprintf("value size, unit is byte, default is %s", DefaultSize),
					Default: DefaultSize,
				},
				&spec.ExpFlag{
					Name:    "duration",
					Desc:    "duration time, default is 10m",
					Default: DefaultDuration,
				},
			},
			ActionExecutor: &CacheHotKeyExecutor{},
			ActionExample: `
# set a key, execute get and set command frequently
blade create redis cache-hot-key --addr 192.168.56.101:6379 --password 123456 --key test1 --thread-count 10 --size 256 --duration 5m
`,
			ActionPrograms:    []string{CacheHotKeyBin},
			ActionCategories:  []string{category.SystemTime},
			ActionProcessHang: true,
		},
	}
}

func (*CacheHotKeyActionCommandSpec) Name() string {
	return ActionCommandName
}

func (*CacheHotKeyActionCommandSpec) Aliases() []string {
	return []string{ShortActionCommandName}
}

func (*CacheHotKeyActionCommandSpec) ShortDesc() string {
	return "Cache Hot Key"
}

func (k *CacheHotKeyActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Set a key, then execute get and set command frequently"
}

func (*CacheHotKeyActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CacheHotKeyExecutor struct {
	channel spec.Channel
}

func (cbe *CacheHotKeyExecutor) Name() string {
	return ActionCommandName
}

func (cbe *CacheHotKeyExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	addrStr := model.ActionFlags["addr"]
	passwordStr := model.ActionFlags["password"]
	keyStr := model.ActionFlags["key"]

	threadCountStr := model.ActionFlags["thread-count"]
	if threadCountStr == "" {
		threadCountStr = DefaultThreadCount
	}

	sizeStr := model.ActionFlags["size"]
	if sizeStr == "" {
		sizeStr = DefaultSize
	}

	durationStr := model.ActionFlags["duration"]
	if durationStr == "" {
		durationStr = DefaultDuration
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return cbe.stop(ctx, addrStr)
	}

	cli := redis.NewClient(&redis.Options{
		Addr:     addrStr,
		Password: passwordStr,
	})

	_, err := cli.Ping(cli.Context()).Result()
	if err != nil {
		errMsg := "redis ping error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	cli.Close()

	return cbe.start(ctx, addrStr, passwordStr, keyStr, threadCountStr, sizeStr, durationStr)
}

func (cbe *CacheHotKeyExecutor) SetChannel(channel spec.Channel) {
	cbe.channel = channel
}

func HotKeyFunc(ctx context.Context, cli *redis.Client, key string, valueForCache string) {
	cli.Set(ctx, key, valueForCache, time.Second*10)
	cli.Get(ctx, key)
}

func (cbe *CacheHotKeyExecutor) start(ctx context.Context, addrStr string, passwordStr string, keyStr string, threadCountStr string, sizeStr string, durationStr string) *spec.Response {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		errMsg := "parse duration-1 error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	threadCount, err := strconv.ParseInt(threadCountStr, 10, 32)
	if err != nil {
		errMsg := "parse thread count error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	size, err := strconv.ParseInt(sizeStr, 10, 32)
	if err != nil {
		errMsg := "parse size error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if size > 500*1024 {
		errMsg := "size can not greater than 500M"
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	endTime := time.Now().Add(duration)

	valueForCache := generateStr(int(size))

	var wg sync.WaitGroup

	for i := 0; i < int(threadCount); i++ {
		wg.Add(1)
		go func() {
			cli := redis.NewClient(&redis.Options{
				Addr:     addrStr,
				Password: passwordStr,
			})
			for {
				if time.Now().Before(endTime) {
					HotKeyFunc(ctx, cli, keyStr, valueForCache)
				} else {
					break
				}
			}

			wg.Done()
		}()
	}

	log.Infof(ctx, "success start hot cache key command")
	wg.Wait()

	return spec.ReturnSuccess("finished hot cache key command success")
}

func (cbe *CacheHotKeyExecutor) stop(ctx context.Context, addrStr string) *spec.Response {
	compile, _ := regexp.Compile(fmt.Sprintf(".*create redis (%s|%s).*%s.*", ActionCommandName, ShortActionCommandName, regexp.QuoteMeta(addrStr)))
	processes, err := process.Processes()
	if err != nil {
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to get processes")
	}

	for _, process := range processes {
		cmdline, err := process.Cmdline()
		if err != nil {
			continue
		}

		if !compile.MatchString(cmdline) {
			continue
		}

		err = process.Kill()
		if err != nil {
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to stop hot cache key command. find the process success, but failed to kill.")
		} else {
			return spec.ReturnSuccess("stop hot cache key command success")
		}
	}

	return spec.ResponseFailWithFlags(spec.ActionNotSupport, "failed to stop hot cache key command. can not find the process success.")
}

func generateStr(size int) string {
	var buffer bytes.Buffer
	for buffer.Len() < size {
		buffer.WriteString("C")
	}

	return buffer.String()
}
