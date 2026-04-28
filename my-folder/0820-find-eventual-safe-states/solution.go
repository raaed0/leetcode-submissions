func eventualSafeNodes(graph [][]int) []int {
    res := []int{}
    safe := make(map[int]bool)

    var dfs func(int) bool
    dfs = func(i int) bool {
        if val, ok := safe[i]; ok {
            return val
        }
        safe[i] = false

        for _, nei := range graph[i] {
            if !dfs(nei) {
                return false
            }
        }
        safe[i] = true
        return safe[i]
    }

    for i := range graph {
        if dfs(i) {
            res = append(res, i)
        }        
    }

    return res
}
