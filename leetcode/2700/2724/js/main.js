/**
 * @param {Array} arr
 * @param {Function} fn
 * @return {Array}
 */
let sortBy = (arr, fn) => arr.sort((a, b) => fn(a) - fn(b));