func change(amount int, coins []int) int {
    memo := make([][]int, len(coins))
    for i := range memo {
        memo[i] = make([]int, amount+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i int, amount int) int {
        if amount == 0 {
            return 1
        }
        if i == len(coins) || amount < 0 {
            return 0
        }

        if memo[i][amount] != -1 {
            return memo[i][amount]
        }

        memo[i][amount] = dfs(i, amount - coins[i]) + dfs(i + 1, amount)
        return memo[i][amount]
    }

    return dfs(0, amount)

}
