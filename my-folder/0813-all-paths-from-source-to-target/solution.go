func allPathsSourceTarget(graph [][]int) [][]int {
    res := [][]int{}
    var dfs func(node int, path []int)
    dfs = func(node int, path []int) {
        if node == len(graph)-1 {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
        }
        for _, nd := range graph[node] {
            dfs(nd, append(path, nd))
        }
    }

    dfs(0, []int{0})
    return res
}
