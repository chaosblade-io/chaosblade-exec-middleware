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
	"strconv"
	"strings"

	"github.com/chaosblade-io/chaosblade-spec-go/log"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"

	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/nginx/parser"
)

const (
	NginxResponseBin           = "chaos_nginxresponse"
	defaultContentType         = "application/json"
	contentTypeHeaderNameUpper = "Content-Type"
	contentTypeHeaderNameLower = "content-type"
	luaCode                    = `local uri = ngx.var.uri;
local path = "%s"
local regex = "%s"

if (path ~= "" and uri == path) or (regex ~= "" and string.match(uri, regex))
then
%s
%s
ngx.exit(%s)
end
`
)

var contentTypeMap = map[string]string{
	"json": "application/json",
	"txt":  "text/plain;charset=utf-8",
}

type ResponseActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewResponseActionSpec() spec.ExpActionCommandSpec {
	return &ResponseActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{
				&spec.ExpFlag{
					Name: "body",
					Desc: "Change response body",
				},
				&spec.ExpFlag{
					Name: "header",
					Desc: "Change response header, you can use ';' to split multiple header kv pairs, such as 'Content-Type=text/plain;Server=mock;'.",
				},
				&spec.ExpFlag{
					Name:    "code",
					Desc:    "Change response code, default 200.",
					Default: "200",
				},
				&spec.ExpFlag{
					Name: "path",
					Desc: "The URI which you will change response on. Note that one of --path and --regex must be set and can't set both at the same time.",
				},
				&spec.ExpFlag{
					Name: "regex",
					Desc: "Change response path through lua regex. Note that one of --path and --regex must be set and can't set both at the same time.",
				},
				&spec.ExpFlag{
					Name: "type",
					Desc: "The new response body type such as json and txt, or you can set Content-Type header to achieve the same function. " +
						"The default type is json.",
					Default: "json",
				},
				&spec.ExpFlag{
					Name:    "server",
					Desc:    "There may be many server blocks in nginx.conf, so which server you want to modify? The default server-id is 0.",
					Default: "0",
				},
				&spec.ExpFlag{
					Name:     "nginx-path",
					Desc:     "The absolute path of nginx",
					Required: true,
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxResponseExecutor{},
			ActionExample: `
# Set /test return body='ok',code=200,type=json
blade create nginx response --path /test --body ok  --nginx-path /usr/local/nginx/sbin/nginx

# Set /test return body='',code=500,type=json
blade create nginx response --path /test --code 500  --nginx-path /usr/local/nginx/sbin/nginx

# Set /test return body='{"a":1}',code=200,type=json
blade create nginx response --path /test --code 200 --body '{"a":1}' --type json --nginx-path /usr/local/nginx/sbin/nginx

# Set /t.* return body='{"a":1}',code=200,type=json, and add header 'Server=mock' to server[0]
blade create nginx response --regex /t.* --code 200 --body '{"a":1}' --header 'Server=mock;' --server 0 --nginx-path /usr/local/nginx/sbin/nginx

# Revert config change to the oldest config file
blade destroy nginx response  --nginx-path /usr/local/nginx/sbin/nginx
`,
			ActionPrograms:   []string{NginxResponseBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*ResponseActionSpec) Name() string {
	return "response"
}

func (*ResponseActionSpec) Aliases() []string {
	return []string{}
}

func (*ResponseActionSpec) ShortDesc() string {
	return "Response experiment"
}

func (d *ResponseActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx response experiment"
}

type NginxResponseExecutor struct {
	channel spec.Channel
}

func (*NginxResponseExecutor) Name() string {
	return "response"
}

func (ng *NginxResponseExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}

	nginxPath := model.ActionFlags["nginx-path"]
	if nginxPath == "" {
		errMsg := "the nginx-path flag is required"
		log.Errorf(ctx, "%s", errMsg)
		return spec.ResponseFailWithFlags(spec.ActionNotSupport, errMsg)
	}
	_, activeFile, _, response := getNginxConfigLocation(ng.channel, ctx, nginxPath)
	if response != nil {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx, nginxPath)
	}
	return ng.start(ctx, activeFile, model)
}

func (ng *NginxResponseExecutor) start(ctx context.Context, activeFile string, model *spec.ExpModel) *spec.Response {
	contentType, response := getContentType(model.ActionFlags["type"])
	if response != nil {
		return response
	}
	newFile, response := setResponse(model, activeFile, contentType, true)
	if response != nil {
		return response
	}

	response = swapNginxConfig(ng.channel, ctx, newFile, model)
	if !response.Success {
		errMsg := response.Err
		if strings.Contains(errMsg, `unknown directive "rewrite_by_lua_block"`) {
			// don't support lua, fallback, output example:
			// nginx: [emerg] unknown directive "content_by_lua_block" in D:\nginx-1.9.9/conf/nginx.conf:43
			// nginx: configuration file D:\nginx-1.9.9/conf/nginx.conf test failed
			newFile, response := setResponse(model, activeFile, contentType, false)
			if response != nil {
				return response
			}
			return swapNginxConfig(ng.channel, ctx, newFile, model)
		}
	}
	return response
}

func setResponse(model *spec.ExpModel, activeFile, contentType string, useLua bool) (string, *spec.Response) {
	path := model.ActionFlags["path"]
	regex := model.ActionFlags["regex"]
	code := model.ActionFlags["code"]
	body := model.ActionFlags["body"]
	header := model.ActionFlags["header"]
	serverId := model.ActionFlags["server"]
	if (regex == "" && path == "") || (regex != "" && path != "") {
		return "", spec.ResponseFailWithFlags(spec.ParameterIllegal, "--path and --regex", "", "--path and --regex can't be empty at the same time")
	}

	config, err := parser.LoadConfig(activeFile)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf("nginx.conf parsing err %s", err))
	}
	server, response := findServerBlock(config, serverId)
	if response != nil {
		return "", response
	}

	newBlock, response := createNewBlock(path, regex, code, body, header, contentType, useLua)
	if response != nil {
		return "", response
	}
	newBlockList := []parser.Block{*newBlock}
	for _, b := range server.Blocks {
		if b.Type != newBlock.Type || b.Header != newBlock.Header {
			newBlockList = append(newBlockList, b)
		}
	}
	server.Blocks = newBlockList

	name := "nginx.chaosblade.tmp.conf"
	err = config.EasyDumpToFile(name)
	if err != nil {
		return "", spec.ReturnFail(spec.OsCmdExecFailed, err.Error())
	}
	return name, nil
}

// The default Content-Type is json
func getContentType(contentTypeKey string) (string, *spec.Response) {
	if contentTypeKey == "" {
		return defaultContentType, nil
	}
	if v, ok := contentTypeMap[contentTypeKey]; ok {
		return v, nil
	}
	support := ""
	for k := range contentTypeMap {
		support += k + ", "
	}
	return "", spec.ResponseFailWithFlags(spec.ParameterInvalid, "--type", contentTypeKey, fmt.Sprintf("--type=%s is not supported, only supports ( %s )", contentTypeKey, support))
}

func findServerBlock(config *parser.Config, id string) (*parser.Block, *spec.Response) {
	if id == "" {
		id = "0"
	}
	serverId, err := strconv.Atoi(id)
	if err != nil {
		return nil, spec.ResponseFailWithFlags(spec.ParameterInvalid, "--server", id, err)
	}

	var http *parser.Block = nil
	for i := 0; i < len(config.Blocks); i++ {
		b := &config.Blocks[i]
		if b.Type == parser.HTTP {
			http = b
			break
		}
	}

	index := 0
	for i := 0; i < len(http.Blocks); i++ {
		b := &http.Blocks[i]
		if b.Type == parser.Server {
			if index == serverId {
				return b, nil
			}
			index++
		}
	}
	if index == 0 {
		return nil, spec.ReturnFail(spec.OsCmdExecFailed, `There is no Server config in nginx.conf`)
	} else {
		return nil, spec.ReturnFail(spec.OsCmdExecFailed, fmt.Sprintf(`Server config not found in nginx.conf, valid serverId : 0 - %d`, index-1))
	}
}

// Create a new Lua block or location block according to the 'useLua' parameter.
func createNewBlock(path, regex, code, body, header, contentType string, useLua bool) (*parser.Block, *spec.Response) {
	block := parser.NewBlock()
	pairs := parseMultipleKvPairs(header)
	if pairs == nil && header != "" {
		return nil, spec.ResponseFailWithFlags(spec.ParameterInvalid, "--header", header, "syntax err")
	}
	if _, err := strconv.Atoi(code); err != nil {
		return nil, spec.ResponseFailWithFlags(spec.ParameterInvalid, "--code", code, "invalid code")
	}

	if useLua {
		block.Type = parser.Lua
		block.Header = "rewrite_by_lua_block"
		headerString := ""
		hasContentType := false
		for _, pair := range pairs {
			headerString += fmt.Sprintf("ngx.header[\"%s\"] = \"%s\"\n", pair[0], pair[1])
			if pair[0] == contentTypeHeaderNameLower ||
				pair[0] == contentTypeHeaderNameUpper {
				hasContentType = true
			}
		}
		if !hasContentType {
			headerString += fmt.Sprintf("ngx.header[\"Content-Type\"] = \"%s\"\n", contentType)
		}
		if body != "" {
			block.Statements = parser.SetStatement(block.Statements,
				fmt.Sprintf(luaCode, path, regex, headerString, fmt.Sprintf("ngx.say('%s')", body), code), "", true)
		} else {
			block.Statements = parser.SetStatement(block.Statements,
				fmt.Sprintf(luaCode, path, regex, headerString, "", code), "", true)
		}
	} else {
		if regex != "" {
			return nil, spec.ReturnFail(spec.OsCmdExecFailed, "Your nginx don't have lua support, so you cannot change response by --regex")
		}
		block.Type = parser.Location
		block.Header = fmt.Sprintf("%s = %s", block.Type, path) // highest priority
		block.Statements = parser.SetStatement(block.Statements, "default_type", fmt.Sprintf("'%s'", contentType), true)
		for _, pair := range pairs {
			block.Statements = parser.SetStatement(block.Statements, "add_header",
				fmt.Sprintf("%s: %s", pair[0], pair[1]), true)
		}
		block.Statements = parser.SetStatement(block.Statements, "return", fmt.Sprintf("%s '%s'", code, body), true)
	}

	return block, nil
}

func (ng *NginxResponseExecutor) stop(ctx context.Context, nginxPath string) *spec.Response {
	return reloadNginxConfig(ng.channel, ctx, nginxPath)
}

func (ng *NginxResponseExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
