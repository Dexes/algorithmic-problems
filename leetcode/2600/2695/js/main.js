/**
 * @param {number[]} nums
 */
let ArrayWrapper = function (nums) {
    this.nums = nums;
};

ArrayWrapper.prototype.valueOf = function () {
    return this.nums.reduce((sum, current) => sum + current, 0);
}

ArrayWrapper.prototype.toString = function () {
    return JSON.stringify(this.nums);
}
