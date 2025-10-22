//go:build windows
// +build windows

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
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

const (
	configBackupName = "nginx.conf.chaosblade.back"
)

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

// Copy nginx.conf to nginx config path and test it.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf(`copy /Y %s %s`, file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("del %s", tmpFile), "") // ignore response
	if !response.Success || !strings.Contains(response.Result.(string), "successful") {
		return response
	}
	return nil
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	response := channel.Run(ctx, "tasklist", "")
	if response.Success {
		processes := response.Result.(string)
		if strings.Contains(processes, "nginx.exe") {
			return nil
		} else {
			return spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx process")
		}
	} else {
		return response
	}
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	response := channel.Run(ctx, "taskkill /im nginx.exe /F", "")
	if !response.Success {
		return response
	} else {
		return nil
	}
}

func runNginxCommand(channel spec.Channel, ctx context.Context, args string) *spec.Response {
	// find nginx location through NGINX_HOME env variable
	dir := os.Getenv("NGINX_HOME")
	if dir == "" {
		return spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx location, check your NGINX_HOME")
	}
	if args == "" {
		// start the nginx daemon, channel.Run will block the goroutine until getting all standard output.
		// On Windows, executing 'nginx.exe' command through channel.Run will block the goroutine and never end.
		c := make(chan *spec.Response)
		go func() {
			c <- channel.Run(ctx, fmt.Sprintf("cd /d %s && start /b nginx", dir), "")
		}()

		select {
		case response := <-c:
			return response
		case <-time.After(time.Duration(1) * time.Second):
			return spec.Success()
		}
	} else {
		return channel.Run(ctx, fmt.Sprintf("cd /d %s && nginx %s", dir, args), "")
	}
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return channel.Run(ctx, fmt.Sprintf("move /Y %s %s", backup, activeFile), "")
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	cmd := ""
	if util.IsExist(backup) {
		// don't create new backup
		cmd = fmt.Sprintf("copy /Y %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("copy %s %s && copy /Y %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && del %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}

// Start nginx process.
func startNginx(channel spec.Channel, ctx context.Context, nginxPath string) *spec.Response {
	return runNginxCommand(channel, ctx, "")
}

// Find nginx config directory, return dir, activeFile, backup.
func getNginxConfigLocation(channel spec.Channel, ctx context.Context, nginxPath string) (string, string, string, *spec.Response) {
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
	// location may be 'D:\nginx-1.9.9/conf/nginx.conf' on windows..
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
func reloadNginxConfig(channel spec.Channel, ctx context.Context, nginxPath string) *spec.Response {
	_, activeFile, backup, response := getNginxConfigLocation(channel, ctx, nginxPath)
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
	nginxPath := model.ActionFlags["nginx-path"]
	dir, activeFile, backup, response := getNginxConfigLocation(channel, ctx, nginxPath)
	if response != nil {
		return response
	}
	if response := testNginxConfig(channel, ctx, newFile, dir); response != nil {
		return response
	}

	// remove auto generated config
	response = backupConfigFile(channel, ctx, backup, activeFile, newFile, model != nil && model.ActionFlags["mode"] != "file")
	if !response.Success {
		return response
	}

	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config changed")
}
