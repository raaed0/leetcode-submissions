func numRollsToTarget(n int, k int, target int) int {
    MOD := int(1e9 + 7)
    dp := make([][]int, n+1)

    for i := range dp {
        dp[i] = make([]int, target + 1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(n, target int) int
    dfs = func(n, target int) int {
        if n == 0 {
            if target == 0 {
                return 1
            }
            return 0
        }
        if target < 0 {
            return 0
        }

        if dp[n][target] != -1 {
            return dp[n][target]
        }

        res := 0
        for val := 1; val <= k; val++ {
            if target - val >= 0 {
                res = (res + dfs(n-1, target - val)) % MOD
            }
        }
        dp[n][target] = res
        return res
    }

    return dfs(n, target)
}
