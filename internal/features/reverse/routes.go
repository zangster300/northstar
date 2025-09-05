package reverse

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"northstar/internal/features/reverse/pages"
)

func SetupRoutes(router chi.Router) error {
	router.Get("/reverse", func(w http.ResponseWriter, r *http.Request) {
		if err := pages.ReversePage().Render(r.Context(), w); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})

	return nil
}
