/**
 * @param {Function} fn
 * @param {Array} args
 * @param {number} t
 * @return {Function}
 */
let cancellable = function (fn, args, t) {
    let id = setTimeout(fn, t, ...args);
    return () => clearTimeout(id);
};