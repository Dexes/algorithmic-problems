/**
 * @param {object} x
 * @return {object}
 */
let undefinedToNull = function (x) {
    if (x === undefined || x === null) {
        return null
    }

    if (Array.isArray(x)) {
        x.forEach((item, index) => x[index] = undefinedToNull(x[index]))
        return x
    }

    if (typeof x === 'object') {
        for (const [key, value] of Object.entries(x))
            x[key] = undefinedToNull(value)

        return x
    }

    return x
};