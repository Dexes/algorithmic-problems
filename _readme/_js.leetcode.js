let elements = Array.from(document.getElementsByClassName('truncate'));
let result = "";
let difficulties = {'Easy': 0, 'Medium': 1, 'Hard': 2};
let idRE = new RegExp('^\\d+');

for (let i in elements) {
    let titleParent = elements[i];
    if (titleParent.tagName !== 'DIV' || titleParent.has || titleParent.children.length === 0) continue;

    let title = titleParent.children[0];
    if (title.tagName !== 'A') continue;

    let difficulty = titleParent.parentNode.parentNode.parentNode.parentNode.parentNode.children[4].children[0];
    if (difficulty.tagName !== 'SPAN') continue;

    let premium = titleParent.parentNode.children[1];

    let url = title.getAttribute('href')
    if (url[url.length - 1] !== '/') url += '/'

    let id = idRE.exec(titleParent.textContent.trim())[0].padStart(4, '0')

    result += JSON.stringify(id) + ': {'
        + '"name": ' + JSON.stringify(titleParent.textContent.trim()) + ", "
        + '"url": ' + JSON.stringify('https://leetcode.com' + url) + ", "
        + '"difficulty": ' + difficulties[difficulty.textContent] + ", "
        + '"premium": ' + (premium ? "true" : "false")
        + '},\n';
}

console.log(result);