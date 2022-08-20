package nginx

import (
	"context"
	"github.com/chaosblade-io/chaosblade-exec-middleware/exec/category"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
)

const NginxCrashBin = "chaos_nginxcrash"

type CrashActionSpec struct {
	spec.BaseExpActionCommandSpec
}

func NewCrashActionSpec() spec.ExpActionCommandSpec {
	return &CrashActionSpec{
		spec.BaseExpActionCommandSpec{
			ActionMatchers: []spec.ExpFlagSpec{},
			ActionFlags:    []spec.ExpFlagSpec{},
			ActionExecutor: &NginxCrashExecutor{},
			ActionExample: `
# Nginx crash
blade create nginx crash

# Nginx restart
blade destroy nginx crash
`,
			ActionPrograms:   []string{NginxCrashBin},
			ActionCategories: []string{category.Middleware},
		},
	}
}

func (*CrashActionSpec) Name() string {
	return "crash"
}

func (*CrashActionSpec) Aliases() []string {
	return []string{}
}

func (*CrashActionSpec) ShortDesc() string {
	return "Crash experiment"
}

func (d *CrashActionSpec) LongDesc() string {
	if d.ActionLongDesc != "" {
		return d.ActionLongDesc
	}
	return "Nginx crash experiment"
}

type NginxCrashExecutor struct {
	channel spec.Channel
}

func (*NginxCrashExecutor) Name() string {
	return "crash"
}

func (ng *NginxCrashExecutor) Exec(suid string, ctx context.Context, model *spec.ExpModel) *spec.Response {
	if _, ok := spec.IsDestroy(ctx); ok {
		return ng.stop(ctx)
	}
	return ng.start(ctx)
}

func (ng *NginxCrashExecutor) start(ctx context.Context) *spec.Response {
	if response := testNginxExists(ng.channel, ctx); response != nil {
		return response
	}
	if response := killNginx(ng.channel, ctx); response != nil {
		return response
	} else {
		return spec.Success()
	}
}

func (ng *NginxCrashExecutor) stop(ctx context.Context) *spec.Response {
	return startNginx(ng.channel, ctx)
}

func (ng *NginxCrashExecutor) SetChannel(channel spec.Channel) {
	ng.channel = channel
}
