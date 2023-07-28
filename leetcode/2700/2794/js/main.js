/**
 * @param {Array} keys
 * @param {Array} values
 * @return {Object}
 */
let createObject = function(keys, values) {
    let result = {}
    for (let i in keys) {
        if (keys[i] in result) continue
        result[keys[i]] = values[i]
    }

    return result
};