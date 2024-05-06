const global = (1, eval)('this'); // https://stackoverflow.com/questions/9107240/1-evalthis-vs-evalthis-in-javascript
global.global = global;
global.globalThis = global;
global.frames = global;
global.self = global;
global.window = global;
global.navigator = {
  userAgent: 'puppeteer',
}
