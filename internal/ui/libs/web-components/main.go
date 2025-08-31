package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	slog.Info("Bundling Web Components...")
	defer slog.Info("Finished!")

	outDir := "../../static/web-components"
	if err := os.RemoveAll(outDir); err != nil {
		return fmt.Errorf("could not remove output directory '%s': %w", outDir, err)
	}

	result := api.Build(api.BuildOptions{
		EntryPoints: []string{
			"reverse-component/index.ts",
		},
		Outdir:            outDir,
		Bundle:            true,
		Write:             true,
		LogLevel:          api.LogLevelInfo,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Format:            api.FormatESModule,
		Sourcemap:         api.SourceMapLinked,
		Target:            api.ESNext,
	})

	if len(result.Errors) > 0 {
		errs := make([]error, len(result.Errors))
		for i, err := range result.Errors {
			errs[i] = errors.New(err.Text)
		}
		return errors.Join(errs...)
	}

	return nil
}
