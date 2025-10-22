/*
 * Copyright 2025 The ChaosBlade Authors
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
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/go-redis/redis/v8"
)

const CacheLimitBin = "chaos_cacheLimit"

type CacheLimitActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCacheLimitActionSpec() spec.ExpActionCommandSpec {
	return &CacheLimitActionCommandSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "addr",
					Desc: "The address of redis server",
				},
				&spec.ExpFlag{
					Name: "password",
					Desc: "The password of server",
				},
				&spec.ExpFlag{
					Name: "size",
					Desc: "The size of cache",
				},
				&spec.ExpFlag{
					Name: "percent",
					Desc: "The percentage of maxmemory",
				},
			},
			ActionExecutor: &CacheLimitExecutor{},
			ActionExample: `
# set maxmemory to 256M
blade create redis cache-limit --addr 192.168.56.101:6379 --password 123456  --size 256M
`,
			ActionPrograms:   []string{CacheLimitBin},
			ActionCategories: []string{category.SystemTime},
		},
	}
}

func (*CacheLimitActionCommandSpec) Name() string {
	return "cache-limit"
}

func (*CacheLimitActionCommandSpec) Aliases() []string {
	return []string{"cl"}
}

func (*CacheLimitActionCommandSpec) ShortDesc() string {
	return "Cache Memory Limit"
}

func (k *CacheLimitActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Set the maxmemory of Redis"
}

func (*CacheLimitActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CacheLimitExecutor struct {
	channel spec.Channel
}

func (cle *CacheLimitExecutor) Name() string {
	return "cache-limit"
}

func (cle *CacheLimitExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	addrStr := model.ActionFlags["addr"]
	passwordStr := model.ActionFlags["password"]
	sizeStr := model.ActionFlags["size"]
	percentStr := model.ActionFlags["percent"]

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

	if _, ok := spec.IsDestroy(ctx); ok {
		// get the value of origin maxmemory
		originCacheSize, err := cli.Get(cli.Context(), "origin_maxmemory_"+uid).Result()
		if err != nil {
			errMsg := "redis get origin max memory error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
		return cle.stop(ctx, cli, originCacheSize)
	}

	// "maxmemory" is an interface that lists with content similar to "[maxmemory 1024]"
	maxmemory, err := cli.ConfigGet(cli.Context(), "maxmemory").Result()
	if err != nil {
		errMsg := "redis get max memory error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	// get the value of maxmemory
	originCacheSize := fmt.Sprint(maxmemory[1])
	return cle.start(ctx, uid, cli, percentStr, originCacheSize, sizeStr)
}

func (cle *CacheLimitExecutor) SetChannel(channel spec.Channel) {
	cle.channel = channel
}

func (cle *CacheLimitExecutor) stop(ctx context.Context, cli *redis.Client, originCacheSize string) *spec.Response {
	result, err := cli.ConfigSet(cli.Context(), "maxmemory", originCacheSize).Result()
	if err != nil {
		errMsg := "redis set max memory error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if result != STATUSOK {
		errMsg := fmt.Sprintf("redis set max memory error: redis command status is %s", result)
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	return spec.ReturnSuccess("cache memory limit restored")
}

func (cle *CacheLimitExecutor) start(ctx context.Context, uid string, cli *redis.Client, percentStr string, originCacheSize string, sizeStr string) *spec.Response {
	var cacheSize string
	if percentStr != "" {
		percentage, err := strconv.ParseFloat(percentStr[0:len(percentStr)-1], 64)
		if err != nil {
			errMsg := "str parse float error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
		originCacheSize, err := strconv.ParseFloat(originCacheSize, 64)
		if err != nil {
			errMsg := "str parse float error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
		cacheSize = fmt.Sprint(int(math.Floor(originCacheSize / 100.0 * percentage)))
	} else {
		cacheSize = sizeStr
	}

	result, err := cli.ConfigSet(cli.Context(), "maxmemory", cacheSize).Result()
	if err != nil {
		errMsg := "redis set max memory error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if result != STATUSOK {
		errMsg := fmt.Sprintf("redis set max memory error: redis command status is %s", result)
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	originErr := cli.Set(cli.Context(), "origin_maxmemory_"+uid, originCacheSize, 0).Err()
	if originErr != nil {
		errMsg := "redis set origin max memory error: " + originErr.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	return spec.ReturnSuccess("cache memory limit changed")
}
