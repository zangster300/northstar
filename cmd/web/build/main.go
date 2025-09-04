package main

import (
	"errors"
	"flag"
	"log/slog"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

const staticDirectory = "internal/ui/static"

func main() {
	watch := flag.Bool("watch", false, "Enable watcher mode")
	flag.Parse()

	if err := run(*watch); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run(watch bool) error {
	opts := api.BuildOptions{
		EntryPointsAdvanced: []api.EntryPoint{
			{
				InputPath:  "internal/ui/libs/web-components/reverse-component/index.ts",
				OutputPath: "libs/web-components",
			},
			// uncomment this EntryPoint once lit + SortableJS have been installed
			// {
			// 	InputPath:  "internal/ui/libs/lit/src/index.ts",
			// 	OutputPath: "libs/lit-components",
			// },
		},
		Outdir:            staticDirectory,
		Bundle:            true,
		Write:             true,
		LogLevel:          api.LogLevelInfo,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Format:            api.FormatESModule,
		Sourcemap:         api.SourceMapLinked,
		Target:            api.ESNext,
	}

	if watch {
		slog.Info("Watching...")
		ctx, err := api.Context(opts)
		if err != nil {
			return err
		}

		if err := ctx.Watch(api.WatchOptions{}); err != nil {
			return err
		}

		<-make(chan struct{})
		return nil
	}

	slog.Info("Building...")

	result := api.Build(opts)

	if len(result.Errors) > 0 {
		errs := make([]error, len(result.Errors))
		for i, err := range result.Errors {
			errs[i] = errors.New(err.Text)
		}
		return errors.Join(errs...)
	}

	return nil
}
