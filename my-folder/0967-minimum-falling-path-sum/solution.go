func minFallingPathSum(matrix [][]int) int {
    M := len(matrix)

    for r := M-2; r >= 0; r-- {
        for c := range M {
            mid := matrix[r+1][c]
            left := math.MaxInt32
            if c > 0 {
                left = matrix[r+1][c-1]
            }
            right := math.MaxInt32
            if c < M-1 {
                right = matrix[r+1][c+1]
            }

            matrix[r][c] = matrix[r][c] + min(left, mid, right)
        }
    }

    res := math.MaxInt32
    for c := 0; c < M; c++ {
        if matrix[0][c] < res {
            res = matrix[0][c]
        } 
    }
    return res
}
