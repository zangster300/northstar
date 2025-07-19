package auth

import (
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router chi.Router) error {
	if err := InitDB("data/users.db"); err != nil {
		return err
	}

	InitSessionStore()

	router.Route("/auth", func(r chi.Router) {
		r.Use(RedirectIfAuthenticated)

		r.Get("/login", HandleLoginPage)
		r.Post("/login", HandleLogin)

		r.Get("/signup", HandleSignupPage)
		r.Post("/signup", HandleSignup)
	})

	router.Post("/auth/logout", HandleLogout)

	return nil
}
