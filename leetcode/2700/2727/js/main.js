/**
 * @param {Object | Array} obj
 * @return {boolean}
 */
let isEmpty = function(obj) {
    for (const _ in obj) return false;
    return true;
};