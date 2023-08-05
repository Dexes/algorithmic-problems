/**
 * @param {Function} callback
 * @param {Object} context
 */
Array.prototype.forEach = function (callback, context) {
    for (let i = 0; i < this.length; i++) {
        callback.call(context, this[i], i, this)
    }
}