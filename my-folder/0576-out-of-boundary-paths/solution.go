func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    MOD := int(1e9) + 7
    memo := make([][][]int, m)

    for r := range memo {
        memo[r] = make([][]int, n)
        for c := range memo[r] {
            memo[r][c] = make([]int, maxMove + 1)
            for moves := range memo[r][c] {
                memo[r][c][moves] = -1
            }
        }
    }

    var dfs func(int, int, int) int 
    dfs = func(r, c, moves int) int {
        if r < 0 || r == m || c < 0 || c == n {
            return 1
        }
        if moves == 0 {
            return 0
        }

        if memo[r][c][moves] != -1 {
            return memo[r][c][moves]
        }

        memo[r][c][moves] = ( (dfs(r+1, c, moves-1) + dfs(r-1, c, moves-1)) % MOD + (dfs(r, c+1, moves-1) + dfs(r, c-1, moves-1)) % MOD ) % MOD

        return memo[r][c][moves]
    }

    return dfs(startRow, startColumn, maxMove)

}
