package routes

import (
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-chi/chi/v5"
	"github.com/starfederation/datastar-go/datastar"
	"github.com/zangster300/northstar/internal/ui/pages"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

func setupMonitorRoute(router chi.Router) error {
	router.Get("/monitor", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.MonitorInitial().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
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
				vm, err := mem.VirtualMemory()
				if err != nil {
					slog.Error("unable to get mem stats", slog.String("error", err.Error()))
					return
				}

				memStats := pages.SystemMonitorSignals{
					MemTotal:       humanize.Bytes(vm.Total),
					MemUsed:        humanize.Bytes(vm.Used),
					MemUsedPercent: fmt.Sprintf("%.2f%%", vm.UsedPercent),
				}

				if err := sse.MarshalAndPatchSignals(memStats); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}

			case <-cpuT.C:
				cpuTimes, err := cpu.Times(false)
				if err != nil {
					slog.Error("unable to get cpu stats", slog.String("error", err.Error()))
					return
				}

				cpuStats := pages.SystemMonitorSignals{
					CpuUser:   relativeTime(cpuTimes[0].User),
					CpuSystem: relativeTime(cpuTimes[0].System),
					CpuIdle:   relativeTime(cpuTimes[0].Idle),
				}

				if err := sse.MarshalAndPatchSignals(cpuStats); err != nil {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}
	})

	return nil
}

func relativeTime(totalSeconds float64) string {
	seconds := int64(math.Round(totalSeconds))

	days := seconds / (24 * 3600)
	seconds %= 24 * 3600

	hours := seconds / 3600
	seconds %= 3600

	minutes := seconds / 60
	seconds %= 60

	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 || len(parts) > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 || len(parts) > 0 {
		parts = append(parts, fmt.Sprintf("%dm", minutes))
	}
	parts = append(parts, fmt.Sprintf("%ds", seconds))

	return strings.Join(parts, " ")
}
