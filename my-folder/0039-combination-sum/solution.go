func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}

    var dfs func(i int, curr []int, total int)
    dfs = func(i int, curr []int, total int) {
        if total == target {
            tmp := make([]int, len(curr))
            copy(tmp, curr)
            res = append(res, tmp)
            return
        }

        if i >= len(candidates) || total > target {
            return
        }

        curr = append(curr, candidates[i])
        dfs(i, curr, total+candidates[i])

        curr = curr[:len(curr)-1]
        dfs(i+1, curr, total)
    }

    dfs(0, []int{}, 0)
    return res
}
