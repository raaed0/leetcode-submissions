func countServers(grid [][]int) int {
    ROWS, COLS := len(grid), len(grid[0])
    rowCnt := make([]int, ROWS)
    colCnt := make([]int, COLS)

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == 1 {
                rowCnt[r]++
                colCnt[c]++
            }
        }
    }

    res := 0
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == 1 && max(rowCnt[r], colCnt[c]) > 1 {
                res++
            }
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
