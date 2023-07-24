/**
 * @param {number} target
 * @return {number}
 */
Array.prototype.upperBound = function (target) {
    let l = 0, r = this.length - 1
    while (l < r) {
        let middle = Math.floor(l + (r - l) / 2 + 1)
        this[middle] > target
            ? r = middle - 1
            : l = middle
    }

    return this[r] === target ? r : -1
};