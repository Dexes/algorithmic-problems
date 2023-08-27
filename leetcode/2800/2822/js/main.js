/**
 * @param {Object} obj
 * @return {Object}
 */
let invertObject = function (obj) {
    let result = {}
    for (const [key, value] of Object.entries(obj)) {
        if (value in result) {
            result[value] = push(result[value], key)
        } else {
            result[value] = key
        }
    }

    return result
};

let push = function (array, item) {
    if (Array.isArray(array)) {
        array.push(item)
        return array
    }

    return [array, item]
}