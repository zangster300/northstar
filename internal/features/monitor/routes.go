package monitor

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	router.Get("/", HandleMonitorPage)
	router.Get("/events", HandleMonitorEvents)

	return nil
}
