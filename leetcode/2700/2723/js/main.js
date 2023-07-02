/**
 * @param {Promise} a
 * @param {Promise} b
 * @return {Promise}
 */
let addTwoPromises = async function (a, b) {
    return Promise.all([a, b]).then(([x, y]) => x + y)
};