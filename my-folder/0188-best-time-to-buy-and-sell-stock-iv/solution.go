func maxProfit(k int, prices []int) int {
    n := len(prices)
    dp := make([][][2]int, n+1)
    for i := range dp {
        dp[i] = make([][2]int, k+1)
        for j := range dp[i] {
            dp[i][j] = [2]int {-1,-1}
        }
    }

    var dfs func(int, int, int) int
    dfs = func(i, trx, hold int) int {
        if i >= len(prices) {
            return 0
        }

        if dp[i][trx][hold] != -1 {
            return dp[i][trx][hold]
        }

        skip := dfs(i+1, trx, hold)
        maxprofit := skip

        if trx > 0 && hold == 0 {
            maxprofit = max(skip, -prices[i] + dfs(i+1, trx-1, 1))
        } else if hold == 1 {
            maxprofit = max(skip, prices[i] + dfs(i+1, trx, 0))
        }

        dp[i][trx][hold] = maxprofit
        return maxprofit
    }

    return dfs(0, k, 0)
}
