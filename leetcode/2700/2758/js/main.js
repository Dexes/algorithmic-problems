/**
 * @return {string}
 */
Date.prototype.nextDay = function () {
    const day = 86400000

    const next = new Date(this.getTime() + day)
    const y = next.getFullYear()
    const m = String(next.getMonth() + 1).padStart(2, '0')
    const d = String(next.getDate()).padStart(2, '0')

    return `${y}-${m}-${d}`
}
