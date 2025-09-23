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
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const (
	CacheExpireBin = "chaos_cacheExpire"
	STATUSOK       = "OK"
	OPTIONNX       = "NX"
	OPTIONXX       = "XX"
	OPTIONGT       = "GT"
	OPTIONLT       = "LT"
)

type CacheExpireActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCacheExpireActionSpec() spec.ExpActionCommandSpec {
	return &CacheExpireActionCommandSpec{
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
					Desc: "The key to be set an expiry, default expire all keys",
				},
				&spec.ExpFlag{
					Name: "expiry",
					Desc: `The expiry of the key. An expiry string should be able to be converted to a time duration, such as "5s" or "30m"`,
				},
				&spec.ExpFlag{
					Name: "option",
					Desc: `
The additional options of expiry, only NX, XX, GT, LT supported:
NX -- Set expiry only when the key has no expiry
XX -- Set expiry only when the key has an existing expiry
GT -- Set expiry only when the new expiry is greater than current one
LT -- Set expiry only when the new expiry is less than current one
`,
				},
			},
			ActionExecutor: &CacheExpireExecutor{},
			ActionExample: `
# expire a key
blade create redis cache-expire --addr 192.168.56.101:6379 --password 123456 --key test1 --expiry 1m

# expire all keys only when the new expiry is greater than current one
blade create redis cache-expire --addr 192.168.56.101:6379 --password 123456 --option GT --expiry 1m
`,
			ActionPrograms:   []string{CacheExpireBin},
			ActionCategories: []string{category.SystemTime},
		},
	}
}

func (*CacheExpireActionCommandSpec) Name() string {
	return "cache-expire"
}

func (*CacheExpireActionCommandSpec) Aliases() []string {
	return []string{"ce"}
}

func (*CacheExpireActionCommandSpec) ShortDesc() string {
	return "Cache Expire"
}

func (k *CacheExpireActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Expire the key in Redis"
}

func (*CacheExpireActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type CacheExpireExecutor struct {
	channel spec.Channel
}

func (cee *CacheExpireExecutor) Name() string {
	return "cache-expire"
}

func (cee *CacheExpireExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	addrStr := model.ActionFlags["addr"]
	passwordStr := model.ActionFlags["password"]
	keyStr := model.ActionFlags["key"]
	expiryStr := model.ActionFlags["expiry"]
	optionStr := model.ActionFlags["option"]

	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnSuccess("destroy set expiry success")
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

	return cee.start(ctx, cli, keyStr, expiryStr, optionStr)
}

func (cee *CacheExpireExecutor) SetChannel(channel spec.Channel) {
	cee.channel = channel
}

func ExpireFunc(cli *redis.Client, key string, expiration time.Duration, option string) *redis.BoolCmd {
	switch option {
	case OPTIONNX:
		// NX -- Set expiry only when the key has no expiry
		return cli.ExpireNX(cli.Context(), key, expiration)
	case OPTIONXX:
		// XX -- Set expiry only when the key has an existing expiry
		return cli.ExpireXX(cli.Context(), key, expiration)
	case OPTIONGT:
		// GT -- Set expiry only when the new expiry is greater than current one
		return cli.ExpireGT(cli.Context(), key, expiration)
	case OPTIONLT:
		// LT -- Set expiry only when the new expiry is less than current one
		return cli.ExpireLT(cli.Context(), key, expiration)
	default:
		return cli.Expire(cli.Context(), key, expiration)
	}
}

func (cee *CacheExpireExecutor) start(ctx context.Context, cli *redis.Client, keyStr string, expiryStr string, optionStr string) *spec.Response {
	expiry, err := time.ParseDuration(expiryStr)
	if err != nil {
		errMsg := "parse duration error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	if keyStr == "" {
		// Get all keys from the server
		allKeys, err := cli.Keys(cli.Context(), "*").Result()
		if err != nil {
			errMsg := "redis get all keys error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}

		for _, key := range allKeys {
			result, err := ExpireFunc(cli, key, expiry, optionStr).Result()
			if err != nil {
				errMsg := "redis expire key error: " + err.Error()
				log.Errorf(ctx, "%s", errMsg)
				return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
			}
			if !result {
				errMsg := "redis expire key failed"
				log.Errorf(ctx, "%s", errMsg)
				return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
			}
		}
	} else {
		result, err := ExpireFunc(cli, keyStr, expiry, optionStr).Result()
		if err != nil {
			errMsg := "redis expire key error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
		if !result {
			errMsg := "redis expire key failed"
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
	}

	return spec.ReturnSuccess("set expiry success")
}
