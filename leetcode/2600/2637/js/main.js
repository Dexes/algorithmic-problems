/**
 * @param {Function} fn
 * @param {number} t
 * @return {Function}
 */
let timeLimit = function (fn, t) {
    return async function (...args) {
        let timeout = new Promise((_, reject) => {
            setTimeout(() => reject('Time Limit Exceeded'), t);
        });

        return Promise.race([fn(...args), timeout]);
    }
};