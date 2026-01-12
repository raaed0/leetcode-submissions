function subsets(nums: number[]): number[][] {
    let result: number[][] = []

    let subset: number[] = []

    let dfs = function(i: number) {
        if (i >= nums.length) {
            result.push([...subset])
            return
        }

        subset.push(nums[i])
        dfs(i + 1)

        subset.pop()
        dfs(i + 1)
    }

    dfs(0)
    return result
};
