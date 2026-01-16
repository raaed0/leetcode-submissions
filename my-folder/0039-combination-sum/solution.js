/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function(candidates, target) {
    let total = 0
    let currConbination = []
    let result = []

    function backtrack(i, currConbination, total) {
        if (total > target || i >= candidates.length) {
            return
        }

        if (total === target) {
            result.push([...currConbination])
            return
        }


        currConbination.push(candidates[i])
        backtrack(i, currConbination, total + candidates[i])


        currConbination.pop()
        backtrack(i+1, currConbination, total)
    }

    backtrack(0, currConbination, total)

    return result
};
