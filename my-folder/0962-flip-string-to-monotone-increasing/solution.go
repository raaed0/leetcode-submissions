func minFlipsMonoIncr(s string) int {
    n := len(s)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = []int{-1,-1}
    }

    var dfs func(i, mono int) int
    dfs = func(i, mono int) int {
        if i == n {
            return 0
        }
        if memo[i][mono] != -1 {
            return memo[i][mono]
        }
        if mono == 1 && s[i] == '0' {
            memo[i][mono] = min(1 + dfs(i+1, 0), dfs(i+1, mono))
        } else if mono == 1 && s[i] == '1' {
            memo[i][mono] = min(1 + dfs(i+1, mono), dfs(i+1, 0))
        } else if mono == 0 && s[i] == '1' {
            memo[i][mono] = dfs(i+1, mono)
        } else {
            memo[i][mono] = 1 + dfs(i + 1, mono)
        }

        return memo[i][mono]
    }

    return dfs(0, 1)
}
