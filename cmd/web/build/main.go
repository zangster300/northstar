package main

import (
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"northstar/config"
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
			/*
				uncomment the entrypoint below after running pnpm install in the web/libs/lit directory
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
		Plugins: []api.Plugin{{
			Name: "hotreload",
			Setup: func(build api.PluginBuild) {
				build.OnEnd(func(result *api.BuildResult) (api.OnEndResult, error) {
					slog.Info("build complete", "errors", len(result.Errors), "warnings", len(result.Warnings))
					if watch && len(result.Errors) == 0 {
						http.Get(fmt.Sprintf("http://%s:%s/hotreload", config.Global.Host, config.Global.Port))
					}
					return api.OnEndResult{}, nil
				})
			},
		}},
		Sourcemap: api.SourceMapLinked,
		Target:    api.ESNext,
		Write:     true,
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
