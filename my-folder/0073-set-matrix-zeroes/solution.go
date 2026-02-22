func setZeroes(matrix [][]int)  {
    rowZero := false
    ROWS := len(matrix)
    if ROWS == 0 {
        return
    }
    COLS := len(matrix[0])

    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if matrix[r][c] == 0 {
                matrix[0][c] = 0

                if r > 0 {
                    matrix[r][0] = 0
                }

                if r == 0 {
                    rowZero = true
                }
            }
        }
    }

    for r := 1; r < ROWS; r++ {
        for c:=1; c<COLS; c++ {
            if matrix[0][c] == 0 || matrix[r][0] == 0 {
                matrix[r][c] = 0
            }
        }
    }
    if matrix[0][0] == 0 {
        for r := 0; r < ROWS; r++ {
            matrix[r][0] = 0
        }
    }
    if rowZero {
        for c := 0; c < COLS; c++ {
            matrix[0][c] = 0
        }
    }

}
