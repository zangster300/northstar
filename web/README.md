# Purpose

This directory holds web resources

# Organization

> [!WARNING]  
> If any pathing is updated, make sure to update esbuild [entrypoints](../cmd/web/build/main.go#L28) and pathing in the [`Taskfile.yml` ](../Taskfile.yml)

## Libs

This directory serves as an entrypoint for any custom JS/TS libraries needed for the project to run

Currently it is being used to hold the following:

- [Custom Element](https://developer.mozilla.org/en-US/docs/Web/API/Web_components/Using_custom_elements) [Web Components](./libs/web-components/)
- [LitElement](https://lit.dev/) [Web Components](./libs/lit/src/components/)

## Resources

### Static

This directory contains static web assets

### Styles

This directory is responsible for styling, it is currently setup to use [TailwindCSS](https://tailwindcss.com/)

### `assets.go`

This file adds some useful pathing variables and sets up the embedded static directory using [hashfs](https://github.com/benbjohnson/hashfs)

### `static_dev.go`

When using the `-tags=dev` build tag, this file supplies an http handler function that serves static assets directly without embedding them, and a function for locating them

### `static_prod.go`

When using the `-tags=prod` build tag (or no build tag), this file supplies an http handler function that embeds static assets directly into the binary using [hashfs](https://github.com/benbjohnson/hashfs), and a function for locating them
