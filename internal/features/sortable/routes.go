package sortable

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	return setupSortableRoute(router)
}