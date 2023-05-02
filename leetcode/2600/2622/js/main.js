let TimeLimitedCache = function () {
    this.cache = new Map()
};

/**
 * @param {number} key
 * @param {number} value
 * @param {number} duration in ms
 * @return {boolean} if un-expired key already existed
 */
TimeLimitedCache.prototype.set = function (key, value, duration) {
    let item = this.cache.get(key)
    if (item) clearTimeout(item.timer)

    this.cache.set(key, {value: value, timer: setTimeout(() => delete this.cache.delete(key), duration)})

    return !!item
};

/**
 * @param {number} key
 * @return {number} value associated with key
 */
TimeLimitedCache.prototype.get = function (key) {
    let item = this.cache.get(key)
    return item ? item.value : -1
};

/**
 * @return {number} count of non-expired keys
 */
TimeLimitedCache.prototype.count = function () {
    return this.cache.size
};