func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := make(map[int][]int)
    preOrder := []int{}
    visited := make(map[int]bool)
    cycle := make(map[int]bool)

    for _, p := range prerequisites {
        a, b := p[0], p[1]
        preMap[a] = append(preMap[a], b)
    }

    var dfs func(int) bool
    dfs = func(sub int) bool {
        if cycle[sub] {
            return false
        }
        if visited[sub] {
            return true
        }
        cycle[sub] = true
    
        for _, p := range preMap[sub] {
            if !dfs(p) {
                return false
            }
        }
        preOrder = append(preOrder, sub)
        cycle[sub] = false
        visited[sub] = true
        return true
    }

    for i:=0; i<numCourses; i++ {
        if !dfs(i) {
            return []int{}
        }
    }

    return preOrder
}
