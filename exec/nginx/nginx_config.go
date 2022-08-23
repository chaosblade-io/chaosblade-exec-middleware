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
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/nginx/parser"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"path/filepath"
)

const (
	NginxConfigBin = "chaos_nginxconfig"
	fileMode       = "file"
	cmdMode        = "cmd"
)

type ConfigActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewConfigActionSpec() spec.ExpActionCommandSpec {
	return &ConfigActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name:     "mode",
					Desc:     fmt.Sprintf("The configuration change mode (%s or %s)", fileMode, cmdMode),
					Required: true,
				},
				&spec.ExpFlag{
					Name: "file",
					Desc: "The new nginx.conf file",
				},
				&spec.ExpFlag{
					Name: "block",
					Desc: "The block locator for config modification, use 'global' if you want to modify the global configuration.",
				},
				&spec.ExpFlag{
					Name: "set-config",
					Desc: "Set multiple key-value config paris for specified block",
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxConfigExecutor{},
			ActionExample: `
# Change config file to my.conf
blade create nginx config --mode file --file my.conf

# Change 'server[0]' exposed on port 8899
blade create nginx config --mode cmd --block 'http.server[0]' --set-config='listen=8899'

# Set 'http.server[0].location[0]' proxy_pass to www.baidu.com
blade create nginx config --mode cmd --block 'http.server[0].location[0]' --set-config='proxy_pass=www.baidu.com'

# Revert config change to the oldest config file
blade destroy nginx config
`,
			ActionPrograms:   []string{NginxConfigBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*ConfigActionSpec) Name() string {
	return "config"
}

func (*ConfigActionSpec) Aliases() []string {
	return []string{}
}

func (*ConfigActionSpec) ShortDesc() string {
	return "Config experiment"
}

func (d *ConfigActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx config"
}

type NginxConfigExecutor struct {
	channel spec.Channel
}

func (*NginxConfigExecutor) Name() string {
	return "config"
}

func (ng *NginxConfigExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}

	_, activeFile, _, response := getNginxConfigLocation(ng.channel, ctx)
	if response != nil {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, model)
	}
	return ng.start(ctx, activeFile, model)
}

func (ng *NginxConfigExecutor) start(ctx context.Context, activeFile string, model *spec.ExpModel) *spec.Response {
	var config *parser.Config
	mode := model.ActionFlags["mode"]
	newFile := model.ActionFlags["file"]
	switch mode {
	case fileMode:
		if newFile == "" || !util.IsExist(newFile) || util.IsDir(newFile) {
			return spec.ResponseFailWithFlags(spec.FileNotExist, fmt.Sprintf("config file '%s'", newFile))
		}
		newFile, _ = filepath.Abs(newFile)
	case cmdMode:
		if config == nil {
			config, _ = parser.LoadConfig(activeFile)
		}
		var resp *spec.Response
		newFile, resp = createNewConfig(config, model.ActionFlags["block"], model.ActionFlags["set-config"])
		if resp != nil {
			return resp
		}
	default:
		return spec.ResponseFailWithFlags(spec.ParameterInvalid, "--mode", mode, fmt.Sprintf("--mode must be '%s' or '%s'", fileMode, cmdMode))
	}

	return swapNginxConfig(ng.channel, ctx, newFile, model)
}

func createNewConfig(config *parser.Config, locator string, newKV string) (string, *spec.Response) {
	if locator == "" {
		return "", spec.ResponseFailWithFlags(spec.ParameterInvalid, "--block", locator, "block locator can't be empty")
	}
	pairs := parseMultipleKvPairs(newKV)
	if pairs == nil {
		return "", spec.ResponseFailWithFlags(spec.ParameterInvalid, "--set-config", newKV, "syntax err")
	}
	for _, pair := range pairs {
		err := config.SetStatement(locator, pair[0], pair[1], false)
		if err != nil {
			return "", spec.ResponseFailWithFlags(spec.OsCmdExecFailed, err.Error())
		}
	}
	name := "nginx.chaosblade.tmp.conf"
	err := config.EasyDumpToFile(name)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	return name, nil
}

func (ng *NginxConfigExecutor) stop(ctx context.Context, model *spec.ExpModel) *spec.Response {
	mode := model.ActionFlags["mode"]
	if mode != "" {
		return spec.ResponseFailWithFlags(spec.ParameterInvalid, "--mode", mode, fmt.Sprintf("--mode cannot be %s when destroying Nginx config experiment", mode))
	}
	return reloadNginxConfig(ng.channel, ctx)
}

func (ng *NginxConfigExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
