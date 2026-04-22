func largestLocal(grid [][]int) [][]int {
    N := len(grid)
    res := make([][]int, N-2)
    for i := range res {
        res[i] = make([]int, N-2)
    }

    for i := 0; i < N-2; i++ {
        for j := 0; j < N-2; j++ {
            for r := i; r < i+3; r++ {
                for c := j; c < j+3; c++ {
                    if grid[r][c] > res[i][j] {
                        res[i][j] = grid[r][c]
                    }
                }
            }
        }
    }

    return res
}
