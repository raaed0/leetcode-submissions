func numSquares(n int) int {
    memo := make([]int, n+1)

    for i := range memo {
        memo[i] = -1
    }

    var dfs func(total int) int
    dfs = func(total int) int {
        if total == 0 {
            return 0
        }
        if memo[total] != -1 {
            return memo[total]
        }
        res := total
        for i := 1; i*i <= total; i++ {
            res = min(res, 1 + dfs(total - i*i))
        }

        memo[total] = res
        return memo[total]
        
    }

    return dfs(n)
}
