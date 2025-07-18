package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"northstar/internal/features/counter"
	"northstar/internal/features/monitor"
	"northstar/internal/features/sortable"
	"northstar/internal/features/todo"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: func() slog.Level {
			logLvlStr := os.Getenv("LOG_LEVEL")
			switch logLvlStr {
			case "DEBUG":
				return slog.LevelDebug
			case "WARN":
				return slog.LevelWarn
			case "ERROR":
				return slog.LevelError
			default:
				return slog.LevelInfo
			}
		}(),
	}))
	slog.SetDefault(logger)

	defer slog.Info("stopping server")

	if err := run(context.Background()); err != nil && err != http.ErrServerClosed {
		slog.Error("error running server", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	sctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	getPort := func() string {
		if p, ok := os.LookupEnv("PORT"); ok {
			return p
		}
		return "8080"
	}

	slog.Info("starting server on port :" + getPort())

	eg, egctx := errgroup.WithContext(sctx)
	eg.Go(func() error {
		router := chi.NewMux()

		router.Use(
			middleware.Logger,
			middleware.Recoverer,
		)

		router.Handle("/static/*", static())

		// Setup session store
		sessionStore := sessions.NewCookieStore([]byte("session-secret"))
		sessionStore.MaxAge(int(24 * time.Hour / time.Second))

		// Setup feature routes
		if err := errors.Join(
			todo.SetupRoutes(router, sessionStore),
			counter.SetupRoutes(router, sessionStore),
			monitor.SetupRoutes(router),
			sortable.SetupRoutes(router),
		); err != nil {
			return fmt.Errorf("error setting up routes: %w", err)
		}

		srv := &http.Server{
			Addr:     "0.0.0.0:" + getPort(),
			Handler:  router,
			ErrorLog: slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
		}

		go func() {
			<-egctx.Done()
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Fatalf("error during shutdown: %v", err)
			}
		}()

		return srv.ListenAndServe()
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
