package main

import (
	"errors"
	"flag"
	"log/slog"
	"northstar/web/resources"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

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
				InputPath:  resources.LibsDirectoryPath + "/web-components/reverse-component/index.ts",
				OutputPath: "libs/reverse-component",
			},
			// uncomment after running pnpm install in the web/libs/lit directory
			// esbuild will only be able to find the lit + sortable libraries after doing so
			// {
			// 	InputPath:  resources.LibsDirectoryPath + "/lit/src/index.ts",
			// 	OutputPath: "libs/sortable-example",
			// },
		},
		Outdir:            resources.StaticDirectoryPath,
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
