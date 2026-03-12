func climbStairs(n int) int {
    dp := make([]int, n+1)
    for i:=0; i<=n; i++ {
        dp[i] = -1
    }
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        if dp[i] != -1 {
            return dp[i]
        }
        dp[i] = dfs(i+1) + dfs(i+2)
        return dp[i]
    }

    return dfs(0)
}
