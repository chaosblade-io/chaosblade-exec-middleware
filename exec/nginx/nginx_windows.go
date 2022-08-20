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
	"strings"
	"time"
)

// Copy nginx.conf to nginx config path and test it.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf(`copy /Y %s %s`, file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("del %s", tmpFile), "") //ignore response
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
	//find nginx location through NGINX_HOME env variable
	dir := os.Getenv("NGINX_HOME")
	if dir == "" {
		return spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx location, check your NGINX_HOME")
	}
	if args == "" {
		//start the nginx daemon, channel.Run will block the goroutine until getting all standard output.
		//On Windows, executing 'nginx.exe' command through channel.Run will block the goroutine and never end.
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
		//don't create new backup
		cmd = fmt.Sprintf("copy /Y %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("copy %s %s && copy /Y %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && del %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}
