package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"northstar/config"
	"northstar/nats"
	"northstar/router"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"golang.org/x/sync/errgroup"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: config.Global.LogLevel,
	}))
	slog.SetDefault(logger)

	if err := run(ctx); err != nil && err != http.ErrServerClosed {
		slog.Error("error running server", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	addr := fmt.Sprintf("%s:%s", config.Global.Host, config.Global.Port)
	slog.Info("server started", "addr", addr)
	defer slog.Info("server shutdown complete")

	eg, egctx := errgroup.WithContext(ctx)

	r := chi.NewMux()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
	)

	sessionStore := sessions.NewCookieStore([]byte(config.Global.SessionSecret))
	sessionStore.MaxAge(86400 * 30)
	sessionStore.Options.Path = "/"
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.Secure = false
	sessionStore.Options.SameSite = http.SameSiteLaxMode

	ns, err := nats.SetupNATS(ctx)
	if err != nil {
		return err
	}

	if err := router.SetupRoutes(egctx, r, sessionStore, ns); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
		BaseContext: func(l net.Listener) context.Context {
			return egctx
		},
		ErrorLog: slog.NewLogLogger(
			slog.Default().Handler(),
			slog.LevelError,
		),
	}

	eg.Go(func() error {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
		return nil
	})

	eg.Go(func() error {
		<-egctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		slog.Debug("shutting down server...")

		if err := srv.Shutdown(shutdownCtx); err != nil {
			slog.Error("error during shutdown", "error", err)
			return err
		}

		return nil
	})

	return eg.Wait()
}
