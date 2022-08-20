package nginx

import (
	"context"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"
	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxRestartBin = "chaos_nginxrestart"

type RestartActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewRestartActionSpec() spec.ExpActionCommandSpec {
	return &RestartActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxRestartExecutor{},
			ActionExample: `
# Nginx restart
blade create nginx restart
`,
			ActionPrograms:   []string{NginxRestartBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*RestartActionSpec) Name() string {
	return "restart"
}

func (*RestartActionSpec) Aliases() []string {
	return []string{}
}

func (*RestartActionSpec) ShortDesc() string {
	return "Restart experiment"
}

func (d *RestartActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx restart experiment"
}

type NginxRestartExecutor struct {
	channel spec.Channel
}

func (*NginxRestartExecutor) Name() string {
	return "restart"
}

func (ng *NginxRestartExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if _, ok := spec.IsDestroy(ctx); ok {
		return spec.ReturnFail(spec.OsCmdExecFailed, "cancel 'nginx restart' is meaningless")
	}
	return ng.start(ctx)
}

func (ng *NginxRestartExecutor) start(ctx context.Context) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}
	if response := killNginx(ng.channel, ctx); response != nil {
		return response
	}
	return startNginx(ng.channel, ctx)
}

func (ng *NginxRestartExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
