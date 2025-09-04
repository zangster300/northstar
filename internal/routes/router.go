package routes

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/benbjohnson/hashfs"
	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/starfederation/datastar-go/datastar"
	"github.com/zangster300/northstar/internal/ui"
)

func SetupRoutes(ctx context.Context, router chi.Router) (err error) {

	var hotReloadOnce sync.Once
	router.Get("/reload", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		hotReloadOnce.Do(func() {
			sse.ExecuteScript("window.location.reload()")
		})
		<-r.Context().Done()
	})

	router.Handle("/static/*", hashfs.FileServer(ui.StaticSys))

	natsPort, err := getFreeNatsPort()
	if err != nil {
		return fmt.Errorf("error obtaining NATS port: %w", err)
	}

	ns, err := embeddednats.New(ctx, embeddednats.WithNATSServerOptions(&natsserver.Options{
		JetStream: true,
		NoSigs:    true,
		Port:      natsPort,
		StoreDir:  "data/nats",
	}))

	if err != nil {
		return fmt.Errorf("error creating embedded nats server: %w", err)
	}

	ns.WaitForServer()
	slog.Info("NATS started", "port", natsPort)

	sessionStore := sessions.NewCookieStore([]byte("session-secret"))
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

	if err := errors.Join(
		setupIndexRoute(router, sessionStore, ns),
		setupCounterRoute(router, sessionStore),
		setupMonitorRoute(router),
		setupSortableRoute(router),
		setupReverseRoute(router),
	); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}

func getFreeNatsPort() (int, error) {
	if p, ok := os.LookupEnv("NATS_PORT"); ok {
		natsPort, err := strconv.Atoi(p)
		if err != nil {
			return 0, fmt.Errorf("error parsing NATS_PORT: %w", err)
		}
		if isPortFree(natsPort) {
			return natsPort, nil
		}
	}
	return toolbelt.FreePort()
}

func isPortFree(port int) bool {
	address := fmt.Sprintf(":%d", port)

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}

	if err := ln.Close(); err != nil {
		return false
	}

	return true
}
