//go:build !dev
// +build !dev

package main

import (
	"log/slog"
	"net/http"

	"northstar/internal/features/common"

	hashFS "github.com/benbjohnson/hashfs"
)

func static() http.Handler {
	slog.Debug("static assets are embedded")
	return hashFS.FileServer(common.StaticAssets)
}
