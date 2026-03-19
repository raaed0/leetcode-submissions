func minSteps(n int) int {
    if n == 1 {
        return 0
    }

    memo := make(map[[2]int]int, n+1)

    var dfs func(int, int)int
    dfs = func(i, j int) int {
        if i == n {
            return 0
        }

        if i > n {
            return 1000
        }

        key := [2]int{i, j}
        if val, ok := memo[key]; ok {
            return val
        }

        memo[key] = min(1 + dfs(i + j, j), 2 + dfs(2 * i, i))
        return memo[key]
    }

    return 1 + dfs(1, 1)
}
