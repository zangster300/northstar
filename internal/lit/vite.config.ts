import fs from 'fs'
import { resolve } from 'path'
import { defineConfig, Plugin } from 'vite'
import pc from 'picocolors'
import dts from 'vite-plugin-dts'

const bundleOutputPath = 'dist'

const copyBuildPlugin = (path: string): Plugin => {
  return {
    name: 'copy-build',
    closeBundle: async () => {
      console.log(`${pc.cyan(`[vite:copyBuild]`)} ${pc.green(`copying bundled files to`)} ${pc.yellow(path)}`)
      const srcPath = resolve(__dirname, bundleOutputPath)
      const destPath = resolve(__dirname, path)
      fs.cp(srcPath, destPath, { recursive: true }, (err) => {
        if (err) {
          console.error(err)
          return
        }
      })
    },
  }
}

export default defineConfig(({ mode }) => {
  return {
    plugins: [dts({ rollupTypes: true, tsconfigPath: './tsconfig.json' }), copyBuildPlugin('../../static/.')],
    build: {
      target: 'esnext',
      minify: 'esbuild',
      outDir: bundleOutputPath,
      sourcemap: mode !== 'production',
      lib: {
        entry: resolve(__dirname, 'src/index.ts'),
        formats: ['es'],
      },
      rollupOptions: {
        output: {
          dir: 'dist',
        },
      },
    },
  }
})
