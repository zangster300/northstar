//go:build !dev
// +build !dev

package main

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"

	hashFS "github.com/benbjohnson/hashfs"
)

//go:embed web/static/*
var staticFS embed.FS
var staticRootFS, _ = fs.Sub(staticFS, "web/static")

func static(logger *slog.Logger) http.Handler {
	logger.Debug("static assets are embedded")
	return hashFS.FileServer(staticRootFS)
}
