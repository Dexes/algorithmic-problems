/**
 * @param {number} init
 * @return {{decrement: (function(): number), increment: (function(): number), reset: (function(): number)}}
 */
let createCounter = function (init) {
    let current = init;

    return {
        increment: () => ++current,
        decrement: () => --current,
        reset: () => current = init,
    }
};