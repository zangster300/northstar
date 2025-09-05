package resources

import (
	"embed"

	"github.com/benbjohnson/hashfs"
)

var (
	LibsDirectoryPath   = "web/libs"
	StylesDirectoryPath = "web/resources/styles"
	StaticDirectoryPath = "web/resources/static"
)

var (
	//go:embed static
	StaticDirectory embed.FS
	StaticSys       = hashfs.NewFS(StaticDirectory)
)

func StaticPath(path string) string {
	return "/" + StaticSys.HashName("static/"+path)
}
