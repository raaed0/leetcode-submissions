function combinationSum(candidates: number[], target: number): number[][] {
    let result: number[][] = []

    let dfs = function(i: number, list: number[], total: number) {
        if (total === target) {
            result.push([...list])
            return
        }

        if (i >= candidates.length || total > target)
            return

        list.push(candidates[i])
        dfs(i, list, total + candidates[i])

        list.pop()
        dfs(i+1, list, total)
    }

    dfs(0, [], 0)
    return result
};
