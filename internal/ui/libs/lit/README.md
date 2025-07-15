# Lit

This directory holds an example [library](https://vite.dev/guide/build.html#library-mode) setup for bundling web components driven by `Datastar`

They are powered by [lit](https://lit.dev/)

# Setup

1. Install Dependencies

```shell
pnpm install
```

2. Build and bundle

```shell
pnpm build
```

- A plugin has been included in [`vite.config.ts`](./vite.config.ts) that automatically copies bundled files to the [`static`](../../static/) directory

# Development

```shell
pnpm build:dev
```
