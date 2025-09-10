# Purpose

This directory holds web resources for an application

# Organization

> [!WARNING]  
> If any pathing is updated, make sure to update esbuild [entrypoints](../cmd/web/build/main.go#L25) and pathing in the [`Taskfile.yml` ](../Taskfile.yml)

## Libs

This directory serves as an entrypoint for any custom JS/TS libraries needed for the project to run

Currently it is being used to hold the following:

- [Custom Elements](./libs/web-components/)
- Web Components powered by [lit](https://lit.dev/)

## Resources

### Static

This directory will be [served by the web server](../internal/router.go#L42)

### Styles

This directory is responsible for styling, it is currently setup to use [TailwindCSS](https://tailwindcss.com/)

### `assets.go`

This file adds pathing variables and sets up the embedded static directory using [hashfs](https://github.com/benbjohnson/hashfs)
