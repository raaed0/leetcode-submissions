func lastStoneWeightII(stones []int) int {
    sum := 0
    for _, v := range stones {
        sum += v
    }

    target := sum / 2

    memo := make(map[[2]int]int)
    var dfs func(int, int) int
    dfs = func(i, total int) int {
        if i == len(stones) || total >= target {
            return abs(total - (sum - total))
        }
        key := [2]int{i, total}
        if val, ok := memo[key]; ok {
            return val
        }

        memo[key] = min(dfs(i + 1, total), dfs(i + 1, total + stones[i]))
        return memo[key]
    }

    return dfs(0, 0)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
