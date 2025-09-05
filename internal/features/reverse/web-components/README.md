# Reverse Component

A vanilla web component that reverses text input, powered by Datastar

## Component

- **File**: [`reverse-component.ts`](./reverse-component.ts)
- **Custom Element**: `<reverse-component>`
- **Output**: `internal/ui/static/web-components/reverse-component.js`

## Setup

1. Install dependencies

```shell
go mod tidy
```

2. Build

```shell
go run cmd/web/build/main.go
```

The component is automatically included in the build process and compiles to `internal/ui/static/web-components/reverse-component.js`.

## Usage

See the implementation in [`internal/features/reverse/pages/reverse.templ`](../pages/reverse.templ) for an example of how this component is used.