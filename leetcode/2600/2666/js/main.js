/**
 * @param {Function} fn
 * @return {Function}
 */
let once = function (fn) {
    return function () {
        if (fn === undefined) return undefined

        let result = fn(...arguments)
        fn = undefined
        return result
    }
};