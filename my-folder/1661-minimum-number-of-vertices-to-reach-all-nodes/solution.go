func findSmallestSetOfVertices(n int, edges [][]int) []int {
    inDegree := make(map[int]int)
    res := []int{}

    for _, edge := range edges {
        inDegree[edge[1]]++
    }

    for i := 0; i < n; i++ {
        if _, ok := inDegree[i]; !ok {
            res = append(res, i)
        }
    }

    return res
}
