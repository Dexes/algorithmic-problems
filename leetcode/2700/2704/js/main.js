/**
 * @param {string} actual
 * @return {Object}
 */
let expect = function (actual) {
    return {
        toBe: (expected) => {
            if (expected === actual) return true;
            throw new Error('Not Equal');
        },
        notToBe: (expected) => {
            if (expected !== actual) return true;
            throw new Error('Equal');
        },
    }
};