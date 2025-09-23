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

const NginxCrashBin = "chaos_nginxcrash"

type CrashActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCrashActionSpec() spec.ExpActionCommandSpec {
	return &CrashActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "nginx-path",
					Desc:     "The absolute path of nginx",
					Required: true,
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxCrashExecutor{},
			ActionExample: `
# Nginx crash
blade create nginx crash --nginx-path /usr/local/nginx/sbin/nginx

# Nginx restart
blade destroy nginx crash --nginx-path /usr/local/nginx/sbin/nginx
`,
			ActionPrograms:   []string{NginxCrashBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*CrashActionSpec) Name() string {
	return "crash"
}

func (*CrashActionSpec) Aliases() []string {
	return []string{}
}

func (*CrashActionSpec) ShortDesc() string {
	return "Crash experiment"
}

func (d *CrashActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx crash experiment"
}

type NginxCrashExecutor struct {
	channel spec.Channel
}

func (*NginxCrashExecutor) Name() string {
	return "crash"
}

func (ng *NginxCrashExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	nginxPath := model.ActionFlags["nginx-path"]
	if nginxPath == "" {
		errMsg := "the nginx-path flag is required"
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, nginxPath)
	}
	return ng.start(ctx)
}

func (ng *NginxCrashExecutor) start(ctx context.Context) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}
	if response := killNginx(ng.channel, ctx); response != nil {
		return response
	} else {
		return spec.Success()
	}
}

func (ng *NginxCrashExecutor) stop(ctx context.Context, nginxPath string) *spec.Response {
	return startNginx(ng.channel, ctx, nginxPath)
}

func (ng *NginxCrashExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
