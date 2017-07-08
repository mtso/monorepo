# go react render

[summary]::
Minimal React app to benchmark server-side rendering.  
Simple React app and server-side render script for testing Golang server-side rendering.

## Rendering API

Takes a context JSON object as input:
```js
const context = {
  "state": [object Object],
  "url": "[request url]"
}
```

Outputs a JSON object:
```js
const output = {
  "markup": "[<markup />]",
  "redirect": "[redirect url if returned from context]"
}
```
