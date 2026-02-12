func maximumImportance(n int, roads [][]int) int64 {
    edges := make([]int, n)

    for _, road := range roads {
        edges[road[0]]++
        edges[road[1]]++
    }

    sort.Ints(edges)
    res := 0
    k := 1

    for _, edge := range edges {
        res += edge * k
        k++
    }

    return int64(res)
}
