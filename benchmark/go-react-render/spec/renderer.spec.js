const expect = require('chai').expect;
const renderMarkup = require('../build/renderMarkup.js').default;

describe('server renderer', function() {
  describe('return value', function() {
    it('returns an object', function() {
      const context = {
        state: { name: 'world' },
        url: '/',
      };

      const output = renderMarkup(context);
      expect(output).to.be.an('object');
      expect(output.markup).to.exist;
      expect(output.redirectUrl).to.be.null;
      expect(output.markup).to.be.a('string');
    });
  });
});
