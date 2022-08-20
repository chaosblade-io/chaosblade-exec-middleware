package nginx

import (
	"context"
	"fmt"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/nginx/parser"
	"strconv"
	"strings"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const (
	NginxResponseBin           = "chaos_nginxresponse"
	defaultContentType         = "text/plain;charset=utf-8"
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
	// "html": "text/html;charset=utf-8",
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
					Desc: "change response body",
				},
				&spec.ExpFlag{
					//为了使body有效，会自动设置content type
					Name: "header",
					Desc: "change response header",
				},
				&spec.ExpFlag{
					Name:    "code",
					Desc:    "change response code, default 200",
					Default: "200",
				},
				&spec.ExpFlag{
					Name: "path",
					Desc: "change response path",
				},
				&spec.ExpFlag{
					Name: "regex",
					Desc: "change response path",
				},
				&spec.ExpFlag{
					Name:    "type",
					Desc:    "new response body type",
					Default: "json",
				},
				&spec.ExpFlag{
					Name:    "server",
					Desc:    "which server you want to modify response? default server id is 0",
					Default: "0",
				},
			},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxResponseExecutor{},
			ActionExample: `
# Set /test returns body='ok',code=200,type=json
blade create nginx response --path /test --body ok

# Set /test returns body='',code=500,type=json
blade create nginx response --path /test --code 500

# Set /test returns body='',code=500,type=json
blade create nginx response --path /test --code 200 --body '{"a":1}' --type json

# Set /t.* returns body='{"a":1}',code=200,type=json,add header 'Server=mock'
blade create nginx response --regex /t.* --code 200 --body '{"a":1}' --header 'Server=mock;' --server 0

# Revert config change to the oldest config file
blade destroy nginx response`,
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

	_, activeFile, _, response := getNginxConfigLocation(ng.channel, ctx)
	if response != nil {
		return response
	}

	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx)
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
			//don't support lua, fallback
			//e.g.,nginx: [emerg] unknown directive "content_by_lua_block" in D:\nginx-1.9.9/conf/nginx.conf:43
			//nginx: configuration file D:\nginx-1.9.9/conf/nginx.conf test failed
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
		return "", spec.ReturnFail(spec.ParameterIllegal, "--path and --regex")
	}

	config, err := parser.LoadConfig(activeFile)
	if err != nil {
		return "", spec.ReturnFail(spec.ParameterIllegal, fmt.Sprintf("nginx.conf parsing err %s", err))
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

func (ng *NginxResponseExecutor) stop(ctx context.Context) *spec.Response {
	return reloadNginxConfig(ng.channel, ctx)
}

func (ng *NginxResponseExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}

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
	return "", spec.ResponseFailWithFlags(spec.ParameterInvalid, "--type", contentTypeKey, fmt.Sprintf("--type %s is not supported, only supports ( %s )", contentTypeKey, support))
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

func createNewBlock(path, regex, code, body, header, contentType string, useLua bool) (*parser.Block, *spec.Response) {
	block := parser.NewBlock()
	pairs := parseMultipleKvPairs(header)
	if pairs == nil && header != "" {
		return nil, spec.ResponseFailWithFlags(spec.ParameterInvalid, "--header", header)
	}
	if _, err := strconv.Atoi(code); err != nil {
		return nil, spec.ResponseFailWithFlags(spec.ParameterInvalid, "--code", code)
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
			return nil, spec.ReturnFail(spec.OsCmdExecFailed, "Your nginx don't have lua support, so cannot change response by --regex")
		}
		block.Type = parser.Location
		block.Header = fmt.Sprintf("%s = %s", block.Type, path) //highest priority
		block.Statements = parser.SetStatement(block.Statements, "default_type", fmt.Sprintf("'%s'", contentType), true)
		for _, pair := range pairs {
			block.Statements = parser.SetStatement(block.Statements, "add_header",
				fmt.Sprintf("%s: %s", pair[0], pair[1]), true)
		}
		block.Statements = parser.SetStatement(block.Statements, "return", fmt.Sprintf("%s '%s'", code, body), true)
	}

	return block, nil
}
