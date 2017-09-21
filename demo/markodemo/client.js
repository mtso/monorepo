var hello = require('./app/hello.marko')

hello.renderSync({name: 'Marko'}).appendTo(document.querySelector('#app'))
