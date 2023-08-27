/**
 * @param {Array<function(): Promise>} functions
 * @param {number} ms
 * @return {Array<function(): Promise>}
 */
let delayAll = function (functions, ms) {
    return functions.map(f => () => makePromise(f, ms))
};

/**
 * @param {function(): Promise} f
 * @param {number} ms
 * @returns {Promise}
 */
let makePromise = function (f, ms) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            f()
                .then(result => resolve(result))
                .catch(error => reject(error))
        }, ms)
    })
}