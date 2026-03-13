func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    cache := make([]int, n)
    for i:=0; i<n; i++ {
        cache[i] = -1
    }

    var dfs func(int) int
    dfs = func(i int) int {
        if i >= n {
            return 0
        }
        if cache[i] != -1 {
            return cache[i]
        }
        cache[i] = cost[i] + min(dfs(i+1), dfs(i+2))
        return cache[i]
    }

    return min(dfs(0), dfs(1))
}
