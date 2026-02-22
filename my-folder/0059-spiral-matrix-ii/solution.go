func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }

    r,c := 0, 0
    dr, dc := 0, 1

    for v := 0; v < n*n; v++ {
        matrix[r][c] = v+1

        nextR := (r + dr + n) % n
        nextC := (c + dc + n) % n

        if matrix[nextR][nextC] != 0 {
            dr, dc = dc, -dr
        }

        r += dr
        c += dc
    }

    return matrix
}
