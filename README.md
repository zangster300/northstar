# NORTHSTAR

# Stack

- [Go](https://go.dev/doc/)
- [NATS](https://docs.nats.io/)
- [Datastar](https://github.com/starfederation/datastar)
- [Templ](https://templ.guide/)
  - [Tailwind](https://tailwindcss.com/) x [DaisyUI](https://daisyui.com/)

# Setup

1. Clone this repository

```shell
git clone https://github.com/zangster300/northstar.git
```

2. Install Dependencies

```shell
go mod tidy
```

3. Create 🚀

# Development

Live Reload is setup out of the box - powered by [Air](https://github.com/air-verse/air) + [esbuild](cmd/web/build/main.go)

Use the [live task](./Taskfile.yml#L76) from the [Taskfile](https://taskfile.dev/) to start with live reload setup

```shell
go tool task live
```

Navigate to [`http://localhost:8080`](http://localhost:8080) in your favorite web browser to begin

## Debugging

The [debug task](./Taskfile.yml#L42) will launch [delve](https://github.com/go-delve/delve) to begin a debugging session with your project's binary

```shell
go tool task debug
```

## IDE Support

- [Templ / TailwindCSS Support](https://templ.guide/commands-and-tools/ide-support)

### Visual Studio Code Integration

[Reference](https://code.visualstudio.com/docs/languages/go)

- [launch.json](./.vscode/launch.json)
- [settings.json](./.vscode/settings.json)

a `Debug Main` configuration has been added to the [launch.json](./.vscode/launch.json)

# Starting the Server

```shell
go tool task run
```

Navigate to [`http://localhost:8080`](http://localhost:8080) in your favorite web browser

# Deployment

## Building an Executable

The `task build` [task](./Taskfile.yml#L33) will assemble and build a binary

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
- [datastar sdk](https://github.com/starfederation/datastar/tree/develop/sdk)
- [templ](https://templ.guide/)

### Embedded NATS

The NATS server that powers the `TODO` application is [embedded into the web server](./cmd/web/main.go#L60)

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

## Web Components x Datastar

[🔗 Vanilla Web Components Setup](./web/libs/web-components/README.md)

[🔗 Lit Web Components Setup](./web/libs/lit/README.md)

## Client

- [datastar](https://www.jsdelivr.com/package/gh/starfederation/datastar)
- [tailwindcss](https://tailwindcss.com/)
- [daisyui](https://daisyui.com/)
- [esbuild](https://esbuild.github.io/)
- [lit](https://lit.dev/)
