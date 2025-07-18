package counter

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store) error {
	getUserValue := GetUserValue(sessionStore)

	router.Get("/counter", HandleCounterPage)
	router.Get("/counter/data", HandleCounterData(getUserValue))
	
	router.Route("/counter/increment", func(incrementRouter chi.Router) {
		incrementRouter.Post("/global", HandleIncrementGlobal)
		incrementRouter.Post("/user", HandleIncrementUser(getUserValue))
	})

	return nil
}