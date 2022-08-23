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
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const configBackupName = "nginx.conf.chaosblade.back"

type NginxCommandSpec struct {
	spec.BaseExpModelCommandSpec
}

func (*NginxCommandSpec) Name() string {
	return "nginx"
}

func (*NginxCommandSpec) ShortDesc() string {
	return "Nginx experiment"
}

func (*NginxCommandSpec) LongDesc() string {
	return "Nginx experiment"
}

func NewNginxCommandSpec() spec.ExpModelCommandSpec {
	return &NginxCommandSpec{
		spec.BaseExpModelCommandSpec{
			ExpActions: []spec.ExpActionCommandSpec{
				NewCrashActionSpec(),
				NewRestartActionSpec(),
				NewConfigActionSpec(),
				NewResponseActionSpec(),
			},
			ExpFlags: []spec.ExpFlagSpec{},
		},
	}
}

// Start nginx process.
func startNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	return runNginxCommand(channel, ctx, "")
}

// Find nginx config directory, return dir, activeFile, backup.
func getNginxConfigLocation(channel spec.Channel, ctx context.Context) (string, string, string, *spec.Response) {
	response := runNginxCommand(channel, ctx, "-t")
	if !response.Success {
		return "", "", "", response
	}
	result := response.Result.(string)
	if !strings.Contains(result, "successful") {
		return "", "", "", spec.ReturnFail(spec.OsCmdExecFailed, `your nginx.conf has something wrong, please run 'nginx -t' to test it.`)
	}
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(result)[1]
	//location may be 'D:\nginx-1.9.9/conf/nginx.conf' on windows..
	location, _ = filepath.Abs(location)
	dir := location[:strings.LastIndex(location, string(os.PathSeparator))+1]
	return dir, location, dir + configBackupName, nil
}

// Parse kv string like 'key=value', 'listen=999;hostname=localhost;'
func parseMultipleKvPairs(newKV string) [][]string {
	if newKV == "" {
		return nil
	}
	var pairs [][]string
	newKV = strings.TrimSpace(newKV)
	newKV = strings.TrimRight(newKV, ";")
	for _, kv := range strings.Split(newKV, ";") {
		arr := strings.Split(strings.TrimSpace(kv), "=")
		if len(arr) != 2 {
			return nil
		}
		k := strings.TrimSpace(arr[0])
		v := strings.TrimSpace(arr[1])
		pairs = append(pairs, []string{k, v})
	}
	return pairs
}

// Reload nginx.conf backup file and send nginx process a reload signal.
func reloadNginxConfig(channel spec.Channel, ctx context.Context) *spec.Response {
	_, activeFile, backup, response := getNginxConfigLocation(channel, ctx)
	if response != nil {
		return response
	}

	if !util.IsExist(backup) || util.IsDir(backup) {
		return spec.ResponseFailWithFlags(spec.FileNotExist, fmt.Sprintf("backup file %s", backup))
	}

	if response := restoreConfigFile(channel, ctx, backup, activeFile); !response.Success {
		return response
	}
	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config restored")
}

// Backup and swap nginx.conf, then send nginx process a reload signal.
func swapNginxConfig(channel spec.Channel, ctx context.Context, newFile string, model *spec.ExpModel) *spec.Response {
	dir, activeFile, backup, response := getNginxConfigLocation(channel, ctx)
	if response != nil {
		return response
	}
	if response := testNginxConfig(channel, ctx, newFile, dir); response != nil {
		return response
	}

	// remove auto generated config
	response = backupConfigFile(channel, ctx, backup, activeFile, newFile, model.ActionFlags["mode"] != fileMode)
	if !response.Success {
		return response
	}

	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config changed")
}
