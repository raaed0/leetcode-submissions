/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var subsets = function(nums) {
    let subset = []
    let result = []

    function backtrack(i) {
        if (i >= nums.length) {
            result.push([...subset])
            return
        }

        subset.push(nums[i])
        backtrack(i+1)

        subset.pop()
        backtrack(i+1)
    }

    backtrack(0)
    return result
};
