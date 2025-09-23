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

	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxRestartBin = "chaos_nginxrestart"

type RestartActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewRestartActionSpec() spec.ExpActionCommandSpec {
	return &RestartActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "nginx-path",
					Desc:     "The absolute path of nginx",
					Required: true,
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxRestartExecutor{},
			ActionExample: `
# Nginx restart
blade create nginx restart  --nginx-path /usr/local/nginx/sbin/nginx
`,
			ActionPrograms:   []string{NginxRestartBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*RestartActionSpec) Name() string {
	return "restart"
}

func (*RestartActionSpec) Aliases() []string {
	return []string{}
}

func (*RestartActionSpec) ShortDesc() string {
	return "Restart experiment"
}

func (d *RestartActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx restart experiment"
}

type NginxRestartExecutor struct {
	channel spec.Channel
}

func (*NginxRestartExecutor) Name() string {
	return "restart"
}

func (ng *NginxRestartExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	nginxPath := model.ActionFlags["nginx-path"]
	if nginxPath == "" {
		errMsg := "the nginx-path flag is required"
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnSuccess("cancel 'nginx restart' is meaningless")
	}
	return ng.start(ctx, nginxPath)
}

func (ng *NginxRestartExecutor) start(ctx context.Context, nginxPath string) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}
	if response := killNginx(ng.channel, ctx); response != nil {
		return response
	}
	return startNginx(ng.channel, ctx, nginxPath)
}

func (ng *NginxRestartExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
