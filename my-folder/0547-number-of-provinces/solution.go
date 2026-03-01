func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        for parent[x] != x {
            parent[x] = parent[parent[x]]
            x = parent[x]
        }
        return x
    }

    union := func(x, y int) {
        rootX, rootY := find(x), find(y)
        if rootX != rootY {
            parent[rootX] = rootY
        }
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if isConnected[i][j] == 1 {
                union(i, j)
            }
        }
    }

    provinces := 0
    for i := 0; i < n; i++ {
        if find(i) == i {
            provinces++
        }
    }

    return provinces
}

