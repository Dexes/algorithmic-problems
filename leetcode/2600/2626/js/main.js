/**
 * @param {number[]} nums
 * @param {Function} fn
 * @param {number} init
 * @return {number}
 */
let reduce = function (nums, fn, init) {
    for (let x of nums) {
        init = fn(init, x)
    }

    return init
};