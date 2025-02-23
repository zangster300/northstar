package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zangster300/northstar/routes"
	"golang.org/x/sync/errgroup"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	getPort := func() string {
		if p, ok := os.LookupEnv("PORT"); ok {
			return p
		}
		return "8080"
	}

	logger.Info("starting server on port :" + getPort())
	defer logger.Info("stopping server")

	if err := run(context.Background(), logger, getPort()); err != nil && err != http.ErrServerClosed {
		logger.Error("error running server", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger, port string) error {
	sctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(sctx)

	eg.Go(func() error {
		router := chi.NewMux()

		router.Use(
			middleware.Logger,
			middleware.Recoverer,
		)

		router.Handle("/static/*", http.StripPrefix("/static/", static(logger)))

		if err := routes.SetupRoutes(ctx, logger, router); err != nil {
			return fmt.Errorf("error setting up routes: %w", err)
		}

		srv := &http.Server{
			Addr:    "0.0.0.0:" + port,
			Handler: router,
		}

		go func() {
			<-ctx.Done()
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
