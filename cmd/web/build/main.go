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
				InputPath:  "internal/features/reverse/web-components/reverse-component.ts",
				OutputPath: "web-components/reverse-component",
			},
			// {
			// 	InputPath:  "internal/features/sortable/web-components/sortable-example.ts",
			// 	OutputPath: "web-components/sortable-example",
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
		NodePaths:         []string{"node_modules"},
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
