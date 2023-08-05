/**
 * @param {number} n
 * @yields {number}
 */
function* factorial(n) {
    if (n === 0) n++

    let current = 1;
    for (let i = 1; i <= n; i++) {
        yield current *= i
    }
}