func longestPalindromeSubseq(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i < 0 || j == n {
            return 0
        }
        if dp[i][j] != -1 {
            return dp[i][j]
        }

        if s[i] == s[j] {
            length := 2
            if i == j {
                length = 1
            }
            dp[i][j] = length + dfs(i-1, j+1)
        } else {
            dp[i][j] = max(dfs(i-1, j), dfs(i, j+1))
        }

        return dp[i][j]
    }

    for i := 0; i < n; i++ {
        dfs(i, i)
        dfs(i, i+1)
    }

    maxLength := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] > maxLength {
                maxLength = dp[i][j]
            }
        }
    }

    return maxLength
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
