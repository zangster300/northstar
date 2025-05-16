//go:build dev
// +build dev

package main

import (
	"log/slog"
	"net/http"
	"os"
)

func static() http.Handler {
	slog.Info("static assets are being served from internal/ui/static/")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.FileServerFS(os.DirFS("internal/ui")).ServeHTTP(w, r)
	})
}
