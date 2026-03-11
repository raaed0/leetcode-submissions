func magnificentSets(n int, edges [][]int) int {
    adj := make(map[int][]int)

    for _, e := range edges {
        adj[e[0]] = append(adj[e[0]], e[1])
        adj[e[1]] = append(adj[e[1]], e[0])
    }

    nodeMaxDepth := make(map[int]int)

    for i := 1; i <= n; i++ {
        queue := []int{i}
        distance := make([]int, n+1)
        distance[i] = 1
        maxDistance := 1
        rootNode := i 

        for len(queue) > 0 {
            currNode := queue[0]
            queue = queue[1:]
            rootNode = min(rootNode, currNode)

            for _, nei := range adj[currNode] {
                if distance[nei] == 0 {
                    distance[nei] = distance[currNode] + 1
                    maxDistance = max(maxDistance, distance[nei])
                    queue = append(queue, nei)
                } else if abs(distance[nei]-distance[currNode]) != 1 {
                    return -1
                }
            }
        }

        nodeMaxDepth[rootNode] = max(nodeMaxDepth[rootNode], maxDistance)
    }

    res := 0
    for _, v := range nodeMaxDepth {
        res += v
    }

    return res
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
