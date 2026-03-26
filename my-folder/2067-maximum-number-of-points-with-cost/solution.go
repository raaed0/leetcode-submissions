func maxPoints(points [][]int) int64 {
    m, n := len(points), len(points[0])
    dp := make([]int64, n)
    for j:=0; j<n; j++ {
        dp[j] = int64(points[0][j])
    }

    for r := 1; r < m; r++ {
        left := make([]int64, n)
        right := make([]int64, n)
        ndp := make([]int64, n)

        left[0] = dp[0]
        for j := 1; j < n; j++ {
            left[j] = max(left[j-1]-1, dp[j])
        }

        right[n-1] = dp[n-1]
        for j := n - 2; j >= 0; j-- {
            right[j] = max(dp[j], right[j+1] - 1)
        }

        for j := 0; j < n; j++ {
            ndp[j] = int64(points[r][j]) + max(left[j], right[j])
        }

        dp = ndp
    }

    var ans int64
    for _, val := range dp {
        if val > ans {
            ans = val
        }
    }

    return ans

}
