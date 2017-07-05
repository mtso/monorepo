const Benchmark = require('benchmark');
const suite = new Benchmark.Suite;
const renderMarkup = require('../build/renderMarkup.js').default;

Benchmark.options.minSamples = 200;

suite
  .add('renderMarkup[url="/"]', function() {
    renderMarkup({ state: { name: 'world' }, url: '/' });
  })
  .on('cycle', function(event) {
    console.log(String(event.target));
  })
  .run({ 'async': true });
