package routes

import (
	"context"
	"errors"
	"time"

	"northstar/internal/features/counter"
	"northstar/internal/features/monitor"
	"northstar/internal/features/sortable"
	"northstar/internal/features/todo"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(ctx context.Context, router chi.Router) (err error) {
	sessionStore := sessions.NewCookieStore([]byte("session-secret"))
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

	if err := errors.Join(
		todo.SetupRoutes(router, sessionStore),
		counter.SetupRoutes(router, sessionStore),
		monitor.SetupRoutes(router),
		sortable.SetupRoutes(router),
	); err != nil {
		return errors.Join(err)
	}

	return nil
}
