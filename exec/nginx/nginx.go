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

//TODO 1.parser
//TODO 2.middleware
//TODO 3.主仓库 ok
//TODO 4.单测,借助于mock channel
//TODO 5.整理一下文档
//TODO 6.编译出最终结果，然后测试 ok
//TODO 7.PR

//FIXME 新的定位形式
//FIXME parser解析lua block
//TODO 整理代码结构
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

func startNginx(channel spec.Channel, ctx context.Context) *spec.Response {
	return runNginxCommand(channel, ctx, "")
}

//dir, activeFile, backup
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

func reloadNginxConfig(channel spec.Channel, ctx context.Context) *spec.Response {
	_, activeFile, backup, response := getNginxConfigLocation(channel, ctx)
	if response != nil {
		return response
	}

	if !util.IsExist(backup) || util.IsDir(backup) {
		return spec.ReturnFail(spec.FileNotExist, fmt.Sprintf("backup file %s", backup))
	}

	if response := restoreConfigFile(channel, ctx, backup, activeFile); !response.Success {
		return response
	}
	if response := runNginxCommand(channel, ctx, "-s reload"); !response.Success {
		return response
	}
	return spec.ReturnSuccess("nginx config restored")
}

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
