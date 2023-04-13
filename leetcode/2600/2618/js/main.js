/**
 * @param {Object} object
 * @param {Function} classFunction
 * @return {boolean}
 */
let checkIfInstanceOf = function (object, classFunction) {
    return object !== null
        && object !== undefined
        && typeof classFunction === 'function'
        && Object(object) instanceof classFunction;
};