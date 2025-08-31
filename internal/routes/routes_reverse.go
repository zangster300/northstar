package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zangster300/northstar/internal/ui/pages"
)

func setupReverseRoute(router chi.Router) error {
	router.Get("/reverse", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.ReversePage().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
