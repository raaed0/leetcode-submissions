func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    adjList := make(map[int][]int)

    for _, e := range edges {
        e1, e2 := e[0], e[1]
        adjList[e1] = append(adjList[e1], e2)
        adjList[e2] = append(adjList[e2], e1)
    }

    degree := make(map[int]int)
    leaves := []int{}

    for node, ngh := range adjList {
        degree[node] = len(ngh)
        if len(ngh) == 1 {
            leaves = append(leaves, node)
        }
    }

    for n > 2 {
        newLeaves := []int{}
        
        for _, leaf := range leaves {
            n--
            for _, neighbor := range adjList[leaf] {
                degree[neighbor]--
                if degree[neighbor] == 1 {
                    newLeaves = append(newLeaves, neighbor)
                }
            }
        }
        leaves = newLeaves
    }

    return leaves
}
