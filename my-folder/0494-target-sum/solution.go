func findTargetSumWays(nums []int, target int) int {
    dp := make(map[[2]int]int)

    var dfs func(int, int) int
    dfs = func(i int, sum int) int {
        if i == len(nums) {
            if sum == target {
                return 1
            }
            return 0
        }

        key := [2]int{i, sum}

        if val, ok := dp[key]; ok {
            return val
        }

        dp[key] = dfs(i+1, sum + nums[i]) + dfs(i+1, sum - nums[i])
        return dp[key]
    }

    return dfs(0, 0)
}
