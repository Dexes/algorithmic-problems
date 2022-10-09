let elements = Array.from(document.getElementsByClassName('truncate'));
let result = "";
let difficulties = {'Easy': 0, 'Medium': 1, 'Hard': 2};

for (let i in elements) {
    let titleParent = elements[i];
    if (titleParent.tagName !== 'DIV' || titleParent.has) continue;

    let title = titleParent.children[0];
    if (title.tagName !== 'A') continue;

    let difficulty = titleParent.parentNode.parentNode.parentNode.parentNode.parentNode.children[4].children[0];
    if (difficulty.tagName !== 'SPAN') continue;

    let premium = titleParent.parentNode.children[1];

    title = title.getAttribute('href')
    if (title[title.length - 1] !== '/') title += '/'

    result += JSON.stringify(titleParent.textContent.trim()) + ': {'
        + '"url": ' + JSON.stringify('https://leetcode.com' + title) + ", "
        + '"difficulty": ' + difficulties[difficulty.textContent] + ", "
        + '"premium": ' + (premium ? "true" : "false")
        + '},\n';
}

console.log(result);