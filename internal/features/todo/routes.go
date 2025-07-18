package todo

import (
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, sessionStore sessions.Store) error {
	mvcSession := getMVCSession(sessionStore)

	router.Get("/", HandleIndexPage)

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/todos", func(todosRouter chi.Router) {
			todosRouter.Get("/", HandleTodosSSE(mvcSession))
			todosRouter.Put("/reset", HandleTodosReset(mvcSession))
			todosRouter.Put("/cancel", HandleTodosCancel(mvcSession))
			todosRouter.Put("/mode/{mode}", HandleTodosSetMode(mvcSession))

			todosRouter.Route("/{idx}", func(todoRouter chi.Router) {
				todoRouter.Post("/toggle", HandleTodoToggle(mvcSession))
				todoRouter.Route("/edit", func(editRouter chi.Router) {
					editRouter.Get("/", HandleTodoEditStart(mvcSession))
					editRouter.Put("/", HandleTodoEditSave(mvcSession))
				})
				todoRouter.Delete("/", HandleTodoDelete(mvcSession))
			})
		})
	})

	return nil
}