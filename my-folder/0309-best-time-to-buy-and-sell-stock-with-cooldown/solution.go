
func maxProfit(prices []int) int {
    n := len(prices)
    dp := make([][]int, n+1)

    for i := range dp {
        dp[i] = []int{-1,-1}
    }
    
    var dfs func(int, int) int
    dfs = func(i int, isBuy int) int {
        if i >= len(prices) {
            return 0
        }

        if dp[i][isBuy] != -1 {
            return dp[i][isBuy]
        }

        cooldown := dfs(i+1, isBuy)

        if isBuy == 1 {
            buy := dfs(i+1, 0) - prices[i]
            dp[i][isBuy] = max(buy, cooldown)
        } else {
            sell := dfs(i+2, 1) + prices[i]
            dp[i][isBuy] = max(sell, cooldown)
        }

        return dp[i][isBuy]
    }

    return dfs(0, 1)
}
