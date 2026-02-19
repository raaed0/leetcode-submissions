func restoreMatrix(rowSum []int, colSum []int) [][]int {
    M := len(rowSum)
    N := len(colSum)

    res := make([][]int, M)
    for i:=0; i<M; i++ {
        res[i] = make([]int, N)
    }

    i, j := 0, 0
    for i < M && j<N {
        val := min(rowSum[i], colSum[j])
        rowSum[i] -= val
        colSum[j] -= val

        res[i][j] = val

        if rowSum[i] == 0 {
            i++
        }
        if colSum[j] == 0 {
            j++
        }

    }

    return res
}
