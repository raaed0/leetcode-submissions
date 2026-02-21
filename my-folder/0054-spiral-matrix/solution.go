func spiralOrder(matrix [][]int) []int {
    res := []int{}

    top, bottom := 0, len(matrix)-1
    left, right := 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        for c := left; c <= right; c++ {
            res = append(res, matrix[top][c])
        }
        top++

        for r := top; r <= bottom; r++ {
            res = append(res, matrix[r][right])
        }
        right--

        if top <= bottom {
            for c := right; c >= left; c-- {
                res = append(res, matrix[bottom][c])
            }
            bottom--
        }

        if left <= right {
            for r := bottom; r >= top; r-- {
                res = append(res, matrix[r][left])
            }
            left++
        }
    }

    return res
}
