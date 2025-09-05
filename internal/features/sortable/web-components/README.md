# Sortable Component

A [Lit](https://lit.dev/) web component with drag-and-drop sorting functionality using [SortableJS](https://github.com/SortableJS/Sortable), powered by [Datastar](https://github.com/starfederation/datastar).

## Component

- **File**: [`sortable-example.ts`](./sortable-example.ts)
- **Custom Element**: `<sortable-example>`
- **Dependencies**: lit, sortablejs
- **Output**: `internal/ui/static/web-components/sortable-component.js`

## Setup

1. Install dependencies. This component requires Lit and SortableJS, which is why it's commented out by default in the build script.

```shell
pnpm install
```

2. Enable the component in the build script

Uncomment [these lines](../../../../cmd/web/build/main.go#L31-L34) in [`cmd/web/build/main.go`](../../../../cmd/web/build/main.go):

```go
{
    InputPath:  "internal/features/sortable/web-components/sortable-example.ts",
    OutputPath: "web-components/sortable-example",
},
```

3. Build

```shell
go run cmd/web/build/main.go
```

This compiles the TypeScript component using esbuild and outputs it to `internal/ui/static/web-components/sortable-component.js`.

## Usage

See the implementation in [`internal/features/sortable/pages/sortable.templ`](../pages/sortable.templ) for an example of how this component is used.
