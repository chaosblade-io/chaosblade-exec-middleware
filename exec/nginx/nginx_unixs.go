//go:build !windows
// +build !windows

package nginx

import (
	"context"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"
)

// nginx.conf may have 'include mime.types;' etc.
func testNginxConfig(channel spec.Channel, ctx context.Context, file, dir string) *spec.Response {
	file, _ = filepath.Abs(file)
	tmpFile := fmt.Sprintf("%snginx_chaosblade_temp_%v.conf", dir, time.Now().Unix())
	response := channel.Run(ctx, fmt.Sprintf("cp %s %s", file, tmpFile), "")
	if !response.Success {
		return response
	}
	response = runNginxCommand(channel, ctx, fmt.Sprintf("-t -c %s", tmpFile))
	_ = channel.Run(ctx, fmt.Sprintf("rm %s", tmpFile), "") //ignore response
	if !response.Success || !strings.Contains(response.Result.(string), "successful") {
		return response
	}
	return nil
}

func testNginxExists(channel spec.Channel, ctx context.Context) *spec.Response {
	_, response := getNginxPid(channel, ctx)
	if response != nil {
		return response
	}
	return nil
}

func killNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	commands := []string{"kill"}
	if response, ok := channel.IsAllCommandsAvailable(ctx, commands); !ok {
		return response
	}

	allPid, response := getNginxPid(channel, ctx)
	if response != nil {
		return response
	}
	//kill master process first
	sort.Ints(allPid)
	for _, pid := range allPid {
		response = channel.Run(ctx, fmt.Sprintf("kill -9 %d", pid), "")
		if !response.Success {
			return response
		}
	}
	return nil
}

func runNginxCommand(channel spec.Channel, ctx context.Context, args string) *spec.Response {
	return channel.Run(ctx, "nginx", args)
}

func getNginxPid(channel spec.Channel, ctx context.Context) ([]int, *spec.Response) {
	response := channel.Run(ctx,
		`ps aux | grep -v grep | egrep 'nginx: ' | awk '{print $2}'`, "")
	if !response.Success {
		return []int{}, response
	}
	result := response.Result.(string)
	count := strings.Count(result, "\n")
	if count == 0 {
		return []int{}, spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx process")
	}
	var allPid []int
	for _, s := range strings.Split(strings.Trim(result, "\n"), "\n") {
		pid, err := strconv.Atoi(s)
		if err != nil {
			return []int{}, spec.ReturnFail(spec.OsCmdExecFailed, "cannot find nginx process")
		}
		allPid = append(allPid, pid)
	}

	return allPid, nil
}

func restoreConfigFile(channel spec.Channel, ctx context.Context, backup, activeFile string) *spec.Response {
	return channel.Run(ctx, fmt.Sprintf("mv -f %s %s", backup, activeFile), "")
}

func backupConfigFile(channel spec.Channel, ctx context.Context, backup string, activeFile string, newFile string, remove bool) *spec.Response {
	cmd := ""
	if util.IsExist(backup) {
		//don't create backup
		cmd = fmt.Sprintf("cp -f %s %s", newFile, activeFile)
	} else {
		cmd = fmt.Sprintf("cp %s %s && cp -f %s %s", activeFile, backup, newFile, activeFile)
	}
	if remove {
		cmd += fmt.Sprintf(" && rm %s", newFile)
	}
	return channel.Run(ctx, cmd, "")
}
