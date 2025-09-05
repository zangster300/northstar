package monitor

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	handlers := NewHandlers()

	router.Get("/monitor", handlers.MonitorPage)
	router.Get("/monitor/events", handlers.MonitorEvents)

	return nil
}
