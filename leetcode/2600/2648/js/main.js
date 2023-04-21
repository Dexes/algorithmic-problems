/**
 * @return {Generator<number>}
 */
let fibGenerator = function* () {
    let [current, next] = [0, 1]

    while (true) {
        yield current;
        [current, next] = [next, current + next];
    }
};