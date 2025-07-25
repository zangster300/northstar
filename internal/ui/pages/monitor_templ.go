// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/starfederation/datastar-go/datastar"
	"github.com/zangster300/northstar/internal/ui/layouts"
)

type SystemMonitorSignals struct {
	MemTotal       string `json:"memTotal,omitempty"`
	MemUsed        string `json:"memUsed,omitempty"`
	MemUsedPercent string `json:"memUsedPercent,omitempty"`
	CpuUser        string `json:"cpuUser,omitempty"`
	CpuSystem      string `json:"cpuSystem,omitempty"`
	CpuIdle        string `json:"cpuIdle,omitempty"`
}

func MonitorInitial() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<nav class=\"flex justify-center my-2\"><ul class=\"menu menu-vertical md:menu-horizontal bg-base-200 rounded-box\"><li class=\"hover:text-primary\"><a href=\"/counter\">Counter Example</a></li><li class=\"hover:text-primary\"><a href=\"/\">Todo Example</a></li><li class=\"hover:text-primary\"><a href=\"/sortable\">Sortable Example</a></li></ul></nav><div id=\"container\" data-on-load=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(datastar.GetSSE("/monitor/events"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/ui/pages/monitor.templ`, Line: 28, Col: 52}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "\" class=\"h-screen flex justify-center items-center\" data-signals=\"{memTotal:'', memUsed:'', memUsedPercent:'', cpuUser:'', cpuSystem:'', cpuIdle:''}\"><div class=\"border border-primary rounded flex gap-8 p-8\"><div id=\"mem\" class=\"flex flex-col\"><h1 class=\"text-center pb-2 text-xl\">Memory</h1><p>Total: <span data-text=\"$memTotal\"></span></p><p>Used: <span data-text=\"$memUsed\"></span></p><p>Used (%): <span data-text=\"$memUsedPercent\"></span></p></div><div id=\"cpu\" class=\"flex flex-col\"><h1 class=\"text-center pb-2 text-xl\">CPU</h1><p>User: <span data-text=\"$cpuUser\"></span></p><p>System: <span data-text=\"$cpuSystem\"></span></p><p>Idle: <span data-text=\"$cpuIdle\"></span></p></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layouts.Base("System Monitoring Example").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
