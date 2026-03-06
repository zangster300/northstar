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
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
	"github.com/gorilla/sessions"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		slog.Error("server failure", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: config.Global.LogLevel,
	}))
	slog.SetDefault(logger)

	r := chi.NewMux()
	r.Use(
		httplog.RequestLogger(logger, nil),
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

	eg, egctx := errgroup.WithContext(ctx)

	if err := router.SetupRoutes(egctx, r, sessionStore, ns); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	addr := fmt.Sprintf("%s:%s", config.Global.Host, config.Global.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
		ErrorLog: slog.NewLogLogger(
			slog.Default().Handler(),
			slog.LevelError,
		),
	}

	eg.Go(func() error {
		slog.Info("server started", "addr", srv.Addr)
		err := srv.ListenAndServe()

		if err == nil || err == http.ErrServerClosed {
			return nil
		}

		return fmt.Errorf("server error: %w", err)
	})

	eg.Go(func() error {
		<-egctx.Done()

		slog.Debug("shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		slog.Debug("shutting down server...")

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown error: %w", err)
		}

		return nil
	})

	return eg.Wait()
}
