# NORTHSTAR - A Hypermedia Application Starter Template

# Stack

- [Go](https://go.dev/doc/)
- [NATS](https://docs.nats.io/)
- [Datastar](https://github.com/starfederation/datastar)
- [Templ](https://templ.guide/)
  - [Tailwind](https://tailwindcss.com/) x [DaisyUI](https://daisyui.com/) x [esbuild](https://esbuild.github.io/)

# Setup

1. Clone this repository

```shell
git clone https://github.com/zangster300/northstar.git
```

2. Install Dependencies

```shell
go mod tidy
go run github.com/go-task/task/v3/cmd/task@latest tools
pnpm install
```

3. Create 🚀

# Development

Live Reload is setup out of the box - powered by [Air](https://github.com/air-verse/air) and [templ](https://templ.guide/developer-tools/live-reload-with-other-tools#putting-it-all-together)'s proxy server

Use the [live task](./Taskfile.yml#L105) from the [Taskfile](https://taskfile.dev/) to start the server

```shell
task live
```

Navigate to [`http://localhost:7331`](http://localhost:7331) in your favorite web browser to begin

> [!WARNING]  
> Currently tracking some buggy behavior where the tasks associated with live-reload will stick around, if you are experiencing this utilize `lsof`, `ps` and `kill` to find and remove the associated processes

## Debugging

The [debug task](./Taskfile.yml#L55) will launch [delve](https://github.com/go-delve/delve) to begin a debugging session with your project's binary

```shell
task debug
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
task run
```

Navigate to [`http://localhost:8080`](http://localhost:8080) in your favorite web browser

# Deployment

## Building an Executable

The `task build` [task](./Taskfile.yml#L45) will assemble and build a binary [with static assets embedded](./static_prod.go#L19)

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
- [nats](https://docs.nats.io/)
- [data-star sdk](https://github.com/starfederation/datastar/tree/develop/sdk)
- [templ](https://templ.guide/)

### Embedded NATS

An embedded NATS server that powers the `TODO` application is configured and booted up in the [router.go](./routes/router.go#L27) file

To interface with it, you should install the [nats-cli](https://github.com/nats-io/natscli)

Here are some commands to inspect and make changes to the bucket backing the `TODO` app:

```shell
# list key value buckets
nats kv ls

# list keys in the `todos` bucket
nats kv ls todos

# get the value for [key]
nats kv get --raw todos [key]

# put a value into [key]
nats kv put todos [key] '{"todos":[{"text":"Hello, NATS!","completed":true}],"editingIdx":-1,"mode":0}'
```

> [!IMPORTANT]  
> To see these updates take place in realtime within the `TODO` example, make sure your browser is pointed to the real server and not the templ proxy server!

## Client

- [tailwindcss](https://tailwindcss.com/)
- [daisyui](https://daisyui.com/)
- [esbuild](https://esbuild.github.io/)
- [lit-html](https://lit.dev/)

### Web Components x Datastar

[🔗 Web Components Setup](./web/libs/lit-html/README.md)

## Hypermedia Architecture

- [Hypermedia-Driven Applications](https://htmx.org/essays/hypermedia-driven-applications/)
- [HTMX Sucks](https://htmx.org/essays/htmx-sucks/)
- [HTMX Sucks (kinda)](https://data-star.dev/essays/htmx_sucks)
- [Streams All the Way Down](https://data-star.dev/essays/event_streams_all_the_way_down)
