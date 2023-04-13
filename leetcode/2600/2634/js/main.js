/**
 * @param {number[]} arr
 * @param {Function} fn
 * @return {number[]}
 */
let filter = function (arr, fn) {
    let wIndex = 0;
    for (let i = 0; i < arr.length; i++) {
        if (!fn(arr[i], i)) continue;

        arr[wIndex] = arr[i]
        wIndex++
    }

    return arr.slice(0, wIndex)
};