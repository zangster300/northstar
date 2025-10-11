//go:build dev

package resources

import (
	"log/slog"
	"net/http"
	"os"
)

func Handler() http.Handler {
	slog.Info("static assets are being served directly", "path", StaticDirectoryPath)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		http.StripPrefix("/static/", http.FileServerFS(os.DirFS(StaticDirectoryPath))).ServeHTTP(w, r)
	})
}

func StaticPath(path string) string {
	return "/static/" + path
}
