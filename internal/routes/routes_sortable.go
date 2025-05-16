package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zangster300/northstar/internal/ui/pages"
)

func setupSortableRoute(router chi.Router) error {
	router.Get("/sortable", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.SortableInitial().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
