package nats

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strconv"

	"github.com/delaneyj/toolbelt"
	"github.com/delaneyj/toolbelt/embeddednats"
	natsserver "github.com/nats-io/nats-server/v2/server"
)

func SetupNATS(ctx context.Context) (*embeddednats.Server, error) {
	natsPort, err := getFreeNatsPort()
	if err != nil {
		return nil, fmt.Errorf("error obtaining NATS port: %w", err)
	}

	ns, err := embeddednats.New(ctx, embeddednats.WithNATSServerOptions(&natsserver.Options{
		JetStream: true,
		NoSigs:    true,
		Port:      natsPort,
		StoreDir:  "data/nats",
	}))

	if err != nil {
		return nil, fmt.Errorf("error creating embedded nats server: %w", err)
	}

	ns.WaitForServer()
	slog.Info("NATS started", "port", natsPort)

	return ns, nil
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
