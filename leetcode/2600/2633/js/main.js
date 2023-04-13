/**
 * @param {any} object
 * @return {string}
 */
let jsonStringify = function (object) {
    const type = typeof object;

    if (type === 'string')
        return '"' + object + '"';

    if (object === null || type === 'number' || type === 'boolean')
        return String(object);

    if (Array.isArray(object))
        return '[' + object.map(jsonStringify).join(',') + ']';

    let items = Object.keys(object).map(key => '"' + key + '":' + jsonStringify(object[key]));
    return '{' + items.join(',') + '}';
};
