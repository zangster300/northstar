package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"northstar/config"
	"northstar/web/resources"
	"os"
	"os/signal"
	"syscall"

	"github.com/evanw/esbuild/pkg/api"
	"golang.org/x/sync/errgroup"
)

var (
	watch = false
)

func main() {
	flag.BoolVar(&watch, "watch", watch, "Enable watcher mode")
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer stop()

	if err := run(ctx); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	eg, egctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return build(egctx)
	})

	return eg.Wait()
}

func build(ctx context.Context) error {
	opts := api.BuildOptions{
		EntryPointsAdvanced: []api.EntryPoint{
			{
				InputPath:  resources.LibsDirectoryPath + "/web-components/reverse-component/index.ts",
				OutputPath: "libs/reverse-component",
			},
			/*
				uncomment the entrypoint below after running pnpm install in the resources.LibsDirectoryPath + /lit directory
				esbuild will only be able to find the lit + sortable libraries after doing so
			*/
			// {
			// 	InputPath:  resources.LibsDirectoryPath + "/lit/src/index.ts",
			// 	OutputPath: "libs/sortable-example",
			// },
		},
		Bundle:            true,
		Format:            api.FormatESModule,
		LogLevel:          api.LogLevelInfo,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		MinifyWhitespace:  true,
		Outdir:            resources.StaticDirectoryPath,
		Sourcemap:         api.SourceMapLinked,
		Target:            api.ESNext,
		Write:             true,
	}

	if watch {
		slog.Info("watching...")

		opts.Plugins = append(opts.Plugins, api.Plugin{
			Name: "hotreload",
			Setup: func(build api.PluginBuild) {
				build.OnEnd(func(result *api.BuildResult) (api.OnEndResult, error) {
					slog.Info("build complete", "errors", len(result.Errors), "warnings", len(result.Warnings))
					if len(result.Errors) == 0 {
						http.Get(fmt.Sprintf("http://%s:%s/hotreload", config.Global.Host, config.Global.Port))
					}
					return api.OnEndResult{}, nil
				})
			},
		})

		buildCtx, err := api.Context(opts)
		if err != nil {
			return err
		}
		defer buildCtx.Dispose()

		if err := buildCtx.Watch(api.WatchOptions{}); err != nil {
			return err
		}

		<-ctx.Done()
		return nil
	}

	slog.Info("building...")

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
