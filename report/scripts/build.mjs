import { build } from 'esbuild'

// set to false for testing
const minify = false;

// need to pre-bundle what gets run in puppeteer.evaluate(...)
// because it gets run in a browser context
process.stdout.write("building report-client... ");
await build({
  minify,
  entryPoints: ['report-client/report.mts'],
  bundle: true,
  sourcemap: true,
  platform: 'browser',
  format: 'esm',
  outfile: 'dist/report-client/report.mjs',
  inject: ['report-client/global-shim.js'],
}).then(() => console.log('done'), () => console.log('failed'));

// then bundle the puppeteer code as it gets run in Node context
process.stdout.write("building puppeteer nodejs server... ");
await build({
  minify,
  entryPoints: ['main.ts'],
  bundle: true,
  sourcemap: true,
  platform: 'node',
  format: 'esm',
  outfile: 'dist/main.mjs',
  external: ['./report-client/report.mjs'],
  inject: ['cjs-shim.js'],
}).then(() => console.log('done'), () => console.log('failed'));;
