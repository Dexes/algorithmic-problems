/**
 * @return {Object}
 */
let createInfiniteObject = function () {
    return new Proxy({}, {
        get(_, prop) {
            return () => String(prop);
        }
    });
};
