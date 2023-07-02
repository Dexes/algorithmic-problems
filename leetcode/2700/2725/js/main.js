/**
 * @param {Function} fn
 * @param {Array} args
 * @param {number} t
 * @return {Function}
 */
let cancellable = function (fn, args, t) {
    fn(...args);

    let id = setInterval(fn, t, ...args);
    return () => clearInterval(id);
};