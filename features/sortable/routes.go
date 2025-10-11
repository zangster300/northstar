package sortable

import (
	"net/http"

	"northstar/features/sortable/pages"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	router.Get("/sortable", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.SortablePage().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
