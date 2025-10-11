//go:build !dev

package resources

import (
	"embed"
	"log/slog"
	"net/http"

	"github.com/benbjohnson/hashfs"
)

var (
	//go:embed static
	StaticDirectory embed.FS
	StaticSys       = hashfs.NewFS(StaticDirectory)
)

func Handler() http.Handler {
	slog.Debug("static assets are embedded")
	return hashfs.FileServer(StaticSys)
}

func StaticPath(path string) string {
	return "/" + StaticSys.HashName("static/"+path)
}
