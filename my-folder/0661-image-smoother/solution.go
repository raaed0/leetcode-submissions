func imageSmoother(img [][]int) [][]int {
    ROWS, COLS := len(img), len(img[0])
    res := make([][]int, ROWS)
    for i := range res {
        res[i] = make([]int, COLS)
    }

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            total, count := 0, 0
            for i := r-1; i <= r+1; i++ {
                for j := c-1; j <= c+1; j++ {
                    if i >= 0 && i < ROWS && j >= 0 && j < COLS {
                        total += img[i][j]
                        count++
                    }
                }
            }

            res[r][c] = total / count
        }
    }

    return res
}
