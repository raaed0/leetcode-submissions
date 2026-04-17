/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var backspaceCompare = function(s, t) {
    return solution(s) === solution(t)
};

function solution(str) {
    let result = []
    for (const s of str) {
        if (s === "#") {
            if (result.length > 0) {
                result.pop()
            }
        } else {
            result.push(s)
        }
    }

    return result.join("")
}
