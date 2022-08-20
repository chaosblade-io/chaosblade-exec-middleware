package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-spec-go/channel"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"regexp"
	"strings"
	"testing"
)

const suid = "12345"

func TestPid(t *testing.T) {
	localChannel := channel.LocalChannel{}
	response := localChannel.Run(context.Background(),
		`ps aux | grep -v grep | egrep 'nginx: master' | awk '{print $2}'`, "")
	fmt.Println(response)
}

func TestCmd(t *testing.T) {
	localChannel := channel.LocalChannel{}
	response := localChannel.Run(context.Background(),
		`ps aux | grep -v grep | egrep -o 'nginx: master.*' | egrep -o ' [^ ]*nginx.*'`, "")
	fmt.Println(response)
}

func TestRegex(t *testing.T) {
	regex := regexp.MustCompile("file (.*) test is successful")
	location := regex.FindStringSubmatch(`nginx: the configuration file /etc/nginx/nginx.conf syntax is ok
	nginx: configuration file /etc/nginx/nginx.conf test is successful`)[1]
	location = location[:strings.LastIndex(location, "/")]
	fmt.Println(location)

	response := testNginxExists(channel.NewLocalChannel(), context.Background())
	fmt.Println(response)

	dir, loc, backup, res := getNginxConfigLocation(channel.NewLocalChannel(), context.Background())
	fmt.Println(dir, loc, backup, res)
}

func TestCrash(t *testing.T) {
	executor := NginxCrashExecutor{channel: channel.NewLocalChannel()}
	model := spec.ExpModel{}
	response := executor.Exec(suid, context.Background(), &model)
	fmt.Println(*response)
}

func TestStart(t *testing.T) {
	executor := NginxCrashExecutor{channel: channel.NewLocalChannel()}
	model := spec.ExpModel{}

	response := executor.Exec(suid, context.WithValue(context.Background(), "suid", suid), &model)
	fmt.Println(*response)
}

func TestRestart(t *testing.T) {
	executor := NginxRestartExecutor{channel: channel.NewLocalChannel()}
	model := spec.ExpModel{}
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	// response := executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	// fmt.Println(*response)
}

func TestConfigChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	model.ActionFlags["mode"] = "file"
	model.ActionFlags["file"] = "conf/ok.conf"
	// model.ActionFlags["file"] = "conf/wrong.conf"
	response := executor.Exec("", context.Background(), &model)
	fmt.Println(*response)

	//cancel
	// response = executor.Exec("dsadsad2", context.WithValue(context.Background(), "suid", "dasdsa"), &model)
	// fmt.Println(*response)
}

func TestConfigChangeRevert(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	//cancel
	response := executor.Exec(suid, context.WithValue(context.Background(), "suid", suid), &model)
	fmt.Println(*response)
}

func TestKVChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["list"] = "true"

	model.ActionFlags["mode"] = "cmd"

	model.ActionFlags["set-config"] = "listen=9999"
	model.ActionFlags["block"] = "http.server[0]"

	// model.ActionFlags["set-config"] = "proxy_pass=https://www.taobao.com"
	// model.ActionFlags["block"] = "http.server[0].location[0]"

	response := executor.Exec(suid, context.Background(), &model)
	fmt.Println(response)
}

func TestCancelKVChange(t *testing.T) {
	s := NewConfigActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	// model.ActionFlags["mode"] = "file"

	response := executor.Exec(suid, context.WithValue(context.Background(), "suid", suid), &model)
	fmt.Println(response)
}

func TestChangeResponse(t *testing.T) {
	s := NewResponseActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)
	model.ActionFlags["type"] = "json"
	model.ActionFlags["path"] = "/"
	//model.ActionFlags["regex"] = "/t.*"
	model.ActionFlags["code"] = "500"
	model.ActionFlags["header"] = "Server=mock;"
	model.ActionFlags["body"] = `{"a":1}`
	// model.ActionFlags["server"] = `0`
	// model.ActionFlags["body"] = "hello!"

	response := executor.Exec(suid, context.Background(), &model)
	fmt.Println(response)
}

func TestCancelResponseChange(t *testing.T) {
	s := NewResponseActionSpec()
	executor := s.Executor()
	executor.SetChannel(channel.NewLocalChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	response := executor.Exec(suid, context.WithValue(context.Background(), "suid", suid), &model)
	fmt.Println(response)
}

func TestTmp(t *testing.T) {
	s := NewResponseActionSpec()
	executor := s.Executor()
	executor.SetChannel(setUpMockChannel())
	model := spec.ExpModel{}
	model.ActionFlags = make(map[string]string)

	response := executor.Exec(suid, context.WithValue(context.Background(), "suid", suid), &model)
	fmt.Println(response)
}

func setUpMockChannel() spec.Channel {
	mockChannel := channel.NewMockLocalChannel().(*channel.MockLocalChannel)
	mockChannel.RunFunc = func(ctx context.Context, script, args string) *spec.Response {
		r := regexp.MustCompile(".*nginx\\s+-t.*")
		if r.MatchString(script) {
			return spec.ReturnSuccess(`nginx: the configuration file /usr/local/openresty/nginx/conf/nginx.conf syntax is ok
nginx: configuration file /usr/local/openresty/nginx/conf/nginx.conf test is successful
`)
		}
		return spec.ReturnSuccess("")
	}
	return mockChannel
}
