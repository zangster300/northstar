package index

import (
	"northstar/features/index/services"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(router chi.Router, store sessions.Store, ns *embeddednats.Server) error {
	todoService, err := services.NewTodoService(ns, store)
	if err != nil {
		return err
	}

	handlers := NewHandlers(todoService)

	router.Get("/", handlers.IndexPage)

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Route("/todos", func(todosRouter chi.Router) {
			todosRouter.Get("/", handlers.TodosSSE)
			todosRouter.Put("/reset", handlers.ResetTodos)
			todosRouter.Put("/cancel", handlers.CancelEdit)
			todosRouter.Put("/mode/{mode}", handlers.SetMode)

			todosRouter.Route("/{idx}", func(todoRouter chi.Router) {
				todoRouter.Post("/toggle", handlers.ToggleTodo)
				todoRouter.Route("/edit", func(editRouter chi.Router) {
					editRouter.Get("/", handlers.StartEdit)
					editRouter.Put("/", handlers.SaveEdit)
				})
				todoRouter.Delete("/", handlers.DeleteTodo)
			})
		})
	})

	return nil
}
