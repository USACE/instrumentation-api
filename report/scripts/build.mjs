import { build } from 'esbuild'

await build({
  entryPoints: ['index.ts'],
  bundle: true,
  sourcemap: true,
  minify: false,
  platform: 'node',
  format: 'cjs',
  outfile: 'dist/index.cjs',
  inject: ['shim.js'],
});
