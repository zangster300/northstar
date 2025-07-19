package diary

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	router.Get("/diary", HandleDiaryPage)
	router.Post("/diary/submit", HandleDiarySubmit)
	
	return nil
}