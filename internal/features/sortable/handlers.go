package sortable

import (
	"net/http"

	"northstar/internal/features/sortable/pages"

	"github.com/go-chi/chi/v5"
)

func setupSortableRoute(router chi.Router) error {
	router.Get("/sortable", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.SortableInitial().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
