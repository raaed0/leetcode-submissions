func numMagicSquaresInside(grid [][]int) int {
    rows, cols := len(grid), len(grid[0])
    res := 0
    for r := 0; r < rows-2; r++ {
        for c := 0; c < cols-2; c++ {
            res += isMagic(grid, r, c)
        }
    }

    return res
}

func isMagic(grid [][]int, r, c int) int {
    set := make(map[int]bool)

    for i:=r; i< r+3; i++ {
        for j := c; j < c+3; j++ {
            val := grid[i][j]
            if val < 1 || val > 9 || set[val] {
                return 0
            }
            set[val] = true
        }
    }

    for i := r; i < r+3; i++ {
        sum := 0
        for j := c; j < c+3; j++ {
            sum += grid[i][j]
        }
        if sum != 15 {
            return 0
        }
    }

    for j := c; j < c+3; j++ {
        sum := 0
        for i := r; i < r+3; i++ {
            sum += grid[i][j]
        }
        if sum != 15 {
            return 0
        }
    }

    diag1 := grid[r][c] + grid[r+1][c+1] + grid[r+2][c+2]
    diag2 := grid[r][c+2] + grid[r+1][c+1] + grid[r+2][c]
    if diag1 != 15 || diag2 != 15 {
        return 0
    }

    return 1
}
