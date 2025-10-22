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

// Package redis-----------------------
// @author:  xiejunqiao
// @since:   2024/11/21
// @desc: //TODO
// ----------------------------------------
package redis

import (
	"context"
	"fmt"

	"github.com/chaosblade-io/chaosblade-exec-os/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/go-redis/redis/v8"
)

const ClientsLimitBin = "chaos_clientLimit"

type ClientsLimitActionCommandSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewClientsLimitActionSpec() spec.ExpActionCommandSpec {
	return &ClientsLimitActionCommandSpec{
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
					Name: "count",
					Desc: "The count of clients",
				},
			},
			ActionExecutor: &ClientsLimitExecutor{},
			ActionExample: `
# set maxclients to 100
blade create redis clients-limit --addr 192.168.56.101:6379 --password 123456  --count 100
`,
			ActionPrograms:   []string{ClientsLimitBin},
			ActionCategories: []string{category.SystemTime},
		},
	}
}

func (*ClientsLimitActionCommandSpec) Name() string {
	return "clients-limit"
}

func (*ClientsLimitActionCommandSpec) Aliases() []string {
	return []string{"cl"}
}

func (*ClientsLimitActionCommandSpec) ShortDesc() string {
	return "Clients Limit"
}

func (k *ClientsLimitActionCommandSpec) LongDesc() string {
	if k.ActionLongDesc != "" {
		return k.ActionLongDesc
	}
	return "Set the clients of Redis"
}

func (*ClientsLimitActionCommandSpec) Categories() []string {
	return []string{category.SystemProcess}
}

type ClientsLimitExecutor struct {
	channel spec.Channel
}

func (cle *ClientsLimitExecutor) Name() string {
	return "clients-limit"
}

func (cle *ClientsLimitExecutor) Exec(uid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	addrStr := model.ActionFlags["addr"]
	passwordStr := model.ActionFlags["password"]
	countStr := model.ActionFlags["count"]
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
		originClientSize, err := cli.Get(cli.Context(), "origin_maxclients_"+uid).Result()
		if err != nil {
			errMsg := "redis get origin max clients error: " + err.Error()
			log.Errorf(ctx, "%s", errMsg)
			return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
		}
		return cle.stop(ctx, cli, originClientSize)
	}
	maxClients, err := cli.ConfigGet(cli.Context(), "maxclients").Result()
	if err != nil {
		errMsg := "redis get max clients error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	originClientsCount := fmt.Sprint(maxClients[1])
	return cle.start(ctx, uid, cli, originClientsCount, countStr)
}

func (cle *ClientsLimitExecutor) SetChannel(channel spec.Channel) {
	cle.channel = channel
}

func (cle *ClientsLimitExecutor) stop(ctx context.Context, cli *redis.Client, originClients string) *spec.Response {
	result, err := cli.ConfigSet(cli.Context(), "maxclients", originClients).Result()
	if err != nil {
		errMsg := "redis set max clients error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if result != STATUSOK {
		errMsg := fmt.Sprintf("redis set max clients error: redis command status is %s", result)
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	return spec.ReturnSuccess("clients limit restored")
}

func (cle *ClientsLimitExecutor) start(ctx context.Context, uid string, cli *redis.Client, originClientsCount string, countStr string) *spec.Response {
	result, err := cli.ConfigSet(cli.Context(), "maxclients", countStr).Result()
	if err != nil {
		errMsg := "redis set max clients error: " + err.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if result != STATUSOK {
		errMsg := fmt.Sprintf("redis set max clients error: redis command status is %s", result)
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	originErr := cli.Set(cli.Context(), "origin_maxclients_"+uid, originClientsCount, 0).Err()

	if originErr != nil {
		errMsg := "redis set origin max clients error: " + originErr.Error()
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}

	return spec.ReturnSuccess(uid)
}
