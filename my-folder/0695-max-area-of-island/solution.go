func maxAreaOfIsland(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    set := make(map[[2]int]bool)

    var dfs func(int, int) int
    dfs = func(r, c int) int {
        if r < 0 || r == rows || c < 0 || c == cols || grid[r][c] == 0 || set[[2]int{r,c}] {
            return 0
        } 
        set[[2]int{r,c}] = true

        return 1 + dfs(r + 1, c) + dfs(r - 1, c) + dfs(r, c + 1) + dfs(r, c - 1)
    }

    island := 0

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            island = max(island, dfs(r, c))
        }
    }

    return island
}
