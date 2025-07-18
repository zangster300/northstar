# NORTHSTAR

# Stack

- [Go](https://go.dev/doc/)
- [Datastar](https://github.com/starfederation/datastar)
- [Templ](https://templ.guide/)
  - [Tailwind](https://tailwindcss.com/) x [DaisyUI](https://daisyui.com/)

# Setup

1. Clone this repository

```shell
git clone https://northstar.git
```

2. Install Dependencies

```shell
go mod tidy
```

3. Create 🚀

# Development

Live Reload is setup out of the box - powered by [Air](https://github.com/air-verse/air) and [templ](https://templ.guide/developer-tools/live-reload-with-other-tools#putting-it-all-together)'s proxy server

Use [live task](./Taskfile.yml#L75) from the [Taskfile](https://taskfile.dev/) to start the server

```shell
go tool task live
```

Navigate to [`http://localhost:7331`](http://localhost:7331) in your favorite web browser to begin

## Port Configuration

The default development port is 8080, but can be configured by creating a `.env` file in the project root:

Example `.env` file:

```sh
PORT=4000
```

## Debugging

The [debug task](./Taskfile.yml#L37) will launch [delve](https://github.com/go-delve/delve) to begin a debugging session with your project's binary

```shell
go tool task debug
```

## IDE Support

- [Templ / TailwindCSS Support](https://templ.guide/commands-and-tools/ide-support)

### Visual Studio Code Integration

[Reference](https://code.visualstudio.com/docs/languages/go)

- [launch.json](./.vscode/launch.json)
- [settings.json](./.vscode/settings.json)

a `Debug Main` configuration has been added to the [launch.json](./.vscode/launch.json) file to set breakpoints

# Starting the Server

```shell
go tool task run
```

Navigate to [`http://localhost:8080`](http://localhost:8080) in your favorite web browser

# Deployment

## Building an Executable

The `task build` [task](./Taskfile.yml#29) will assemble and build a binary [with static assets embedded](./cmd/web/static_prod.go#L14)

## Docker

```shell
# build an image
docker build -t northstar:latest .

# run the image in a container
docker run --name northstar -p 8080:9001 northstar:latest
```

[Dockerfile](./Dockerfile)

# Contributing

Completely open to PR's and feature requests

# References

## Server

- [go](https://go.dev/)
- [datastar sdk](https://github.com/starfederation/datastar/tree/develop/sdk)
- [templ](https://templ.guide/)

### In-Memory Storage

The `TODO` application uses in-memory storage with real-time updates powered by Go channels and server-sent events. Todo data is stored per session and automatically synchronized across browser tabs for the same session.

## Web Components x Lit x Datastar

[🔗 Web Components Setup](./internal/ui/libs/lit/README.md)

## Client

- [tailwindcss](https://tailwindcss.com/)
- [daisyui](https://daisyui.com/)
- [esbuild](https://esbuild.github.io/)
- [lit](https://lit.dev/)
