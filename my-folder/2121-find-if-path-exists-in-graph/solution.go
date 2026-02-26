func validPath(n int, edges [][]int, source int, destination int) bool {
    if source == destination {
        return true
    }
    
    adj := make(map[int][]int)
    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }

    level := []int{source}
    visited := make(map[int]bool)
    visited[source] = true

    for len(level) > 0 {
        node := level[0]
        level = level[1:]

        if node == destination {
            return true
        }
        
        for _, ngh := range adj[node] {
            if !visited[ngh] {
                visited[ngh] = true
                level = append(level, ngh)
            }
        }
    }

    return false
}
