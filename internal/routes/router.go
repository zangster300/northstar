package routes

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(ctx context.Context, router chi.Router) (err error) {
	sessionStore := sessions.NewCookieStore([]byte("session-secret"))
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

	if err := errors.Join(
		setupIndexRoute(router, sessionStore),
		setupCounterRoute(router, sessionStore),
		setupMonitorRoute(router),
		setupSortableRoute(router),
	); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}

