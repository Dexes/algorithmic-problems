/**
 * @param {Number} times
 * @return {String}
 */
String.prototype.replicate = function (times) {
    let result = this, tail = ""
    for (; times > 1; times >>= 1) {
        if (times & 1) tail += result
        result += result
    }

    return result + tail
}
