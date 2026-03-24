func minPathSum(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    memo := make([][]int, m)

    for i := range m {
        memo[i] = make([]int, n)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(int, int) int
    dfs = func(i, j int) int {
        if i >= m || j >= n {
            return math.MaxInt32
        }
        if i == m-1 && j == n-1 {
            return grid[i][j]
        }

        if memo[i][j] != -1 {
            return memo[i][j]
        }

        memo[i][j] = grid[i][j] + min(dfs(i + 1, j), dfs(i, j+1))
        return memo[i][j]
    }

    return dfs(0, 0)
}
