package routes

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-chi/chi/v5"
	datastar "github.com/starfederation/datastar/sdk/go"
	"github.com/zangster300/northstar/web/pages"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func setupMonitorRoute(router chi.Router) error {
	router.Get("/monitor", func(w http.ResponseWriter, r *http.Request) {
		pages.MonitorInitial().Render(r.Context(), w)
		return
	})

	router.Get("/monitor/events", func(w http.ResponseWriter, r *http.Request) {

		memT := time.NewTicker(time.Second)
		defer memT.Stop()

		cpuT := time.NewTicker(time.Second)
		defer cpuT.Stop()

		sse := datastar.NewSSE(w, r)
		for {
			select {
			case <-r.Context().Done():
				slog.Debug("client disconnected")
				return

			case <-memT.C:
				m, err := mem.VirtualMemory()
				if err != nil {
					slog.Error("unable to get mem stats", slog.String("error", err.Error()))
					return
				}

				memStats := pages.SystemMonitorSignals{
					MemTotal:       humanize.Bytes(m.Total),
					MemUsed:        humanize.Bytes(m.Used),
					MemUsedPercent: fmt.Sprintf("%.2f%%", m.UsedPercent),
				}

				sse.MarshalAndMergeSignals(memStats)

			case <-cpuT.C:
				c, err := cpu.Times(false)
				if err != nil {
					slog.Error("unable to get cpu stats", slog.String("error", err.Error()))
					return
				}

				cpuStats := pages.SystemMonitorSignals{
					CpuUser:   fmt.Sprintf("%.2f", c[0].User),
					CpuSystem: fmt.Sprintf("%.2f", c[0].System),
					CpuIdle:   fmt.Sprintf("%.2f", c[0].Idle),
				}

				sse.MarshalAndMergeSignals(cpuStats)
			}
		}
	})

	return nil
}
