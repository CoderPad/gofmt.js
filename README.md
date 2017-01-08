This is the `gofmt` tool compiled via [GopherJS](http://www.gopherjs.org) into JavaScript. It's about a megabyte big, and works with Node and the browser (and probably require.js too). The interface is just one function:

```js
var gofmt = require('gofmt.js');
console.log(gofmt("func main() {\nprintln() }"));
// prints:
// func main() {
//     println()
// }
```

If you include `gofmt.js` in the browser, the formatting function is exposed as `window.gofmt`. This function returns false if there was an error, and a string containing the formatted code otherwise.

As of this writing, this was compiled from Go 1.7.4's formatting code.
