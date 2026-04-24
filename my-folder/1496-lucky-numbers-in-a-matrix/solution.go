func luckyNumbers(matrix [][]int) []int {
    res := []int{}
    rows, cols := len(matrix), len(matrix[0])
    max_of_rows_min := math.MinInt
    for r := 0; r < rows; r++ {
        rows_min := math.MaxInt
        for c := 0; c < cols; c++ {
            rows_min = min(matrix[r][c], rows_min)
        }
        max_of_rows_min = max(max_of_rows_min, rows_min)
    }

    for c := 0; c < cols; c++ {
        col_max := matrix[0][c]
        for r := 0; r < rows; r++ {
            col_max = max(col_max, matrix[r][c])
        }
        if col_max == max_of_rows_min {
            res = append(res, col_max)
        }
    }

    return res
}
