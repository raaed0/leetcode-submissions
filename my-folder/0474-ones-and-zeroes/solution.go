func findMaxForm(strs []string, m int, n int) int {
    memo := make(map[[3]int]int)

    var dfs func(int, int, int) int
    dfs = func(i, m, n int) int {
        if i == len(strs) {
            return 0
        }
        key := [3]int{i, m, n}

        if val, ok := memo[key]; ok {
            return val
        }

        mm, nn := strings.Count(strs[i], "0"), strings.Count(strs[i], "1")
        memo[key] = dfs(i + 1, m, n)

        if mm <= m && nn <= n {
            memo[key] = max(memo[key], 1 + dfs(i + 1, m - mm, n - nn))
        }

        return memo[key]
    }

    return dfs(0, m, n)   
}
