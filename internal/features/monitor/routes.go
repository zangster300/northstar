package monitor

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	router.Get("/monitor", HandleMonitorPage)
	router.Get("/monitor/events", HandleMonitorEvents)
	
	return nil
}