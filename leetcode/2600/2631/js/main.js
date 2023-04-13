/**
 * @param {Function} fn
 * @return {Array}
 */
Array.prototype.groupBy = function (fn) {
    let result = {}
    for (let x of this) {
        const key = fn(x);

        if (result[key] === undefined) {
            result[key] = [x];
        } else {
            result[key].push(x);
        }
    }

    return result
};