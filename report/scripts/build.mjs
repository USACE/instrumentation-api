import { build } from 'esbuild'

// set to false for testing
const minify = false;

// need to pre-bundle what gets run in puppeteer.evaluate(...)
// because it gets run in a browser context
process.stdout.write("building report client... ");
await build({
  minify,
  entryPoints: ['src/report.mts'],
  bundle: true,
  sourcemap: true,
  platform: 'browser',
  format: 'esm',
  outfile: 'dist/report.mjs',
}).then(() => console.log('done'), err => { throw new Error(err) });

// then bundle the puppeteer code as it gets run in Node context
process.stdout.write("building puppeteer nodejs server... ");
await build({
  minify,
  entryPoints: ['src/main.ts'],
  bundle: true,
  sourcemap: true,
  platform: 'node',
  format: 'esm',
  outfile: 'dist/main.mjs',
  external: ['./report.mjs'],
  supported: {
    'dynamic-import': true,
  },
  inject: ['src/cjs-shim.js'],
}).then(() => console.log('done'), err => { throw new Error(err) });
