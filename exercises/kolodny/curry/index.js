// Curry-er
// mtso 2017
//
// Usage:
// add = curryer((a, b) => a + b)
// console.log(add(2)(2)) // 4

module.exports = function curryer(callback) {
  return function curry() {
    var args = Array.prototype.slice.call(arguments)

    if (args.length >= callback.length) {
      return callback.apply(null, args)
    } else {
      return curry.bind.apply(curry, [null].concat(args))
    }
  }
}
