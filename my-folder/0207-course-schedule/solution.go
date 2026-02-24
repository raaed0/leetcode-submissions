func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
    visited := make(map[int]bool)

    for _, pre := range prerequisites {
        a, b := pre[0], pre[1]
        preMap[a] = append(preMap[a], b)
    }

    var dfs func(node int) bool
    dfs = func(node int) bool {
        if visited[node] {
            return false
        }
        if len(preMap[node]) == 0 {
            return true
        }
        visited[node] = true

        for _, p := range preMap[node] {
            if !dfs(p) {
                return false
            }
        }

        visited[node] = false
        preMap[node] = []int{}
        return true

    }

    for course := 0; course < numCourses; course++ {
        if !dfs(course) {
            return false
        }
    }
    return true
}
