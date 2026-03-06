func isBipartite(graph [][]int) bool {
    n := len(graph)
    dsu := NewDSU(n)

    for i:=0; i<n; i+=1 {
        for _, v := range graph[i] {
            dsu.Union(graph[i][0], v)

            if dsu.Find(i) == dsu.Find(v) {
                return false
            }
        }
    }

    return true
}

type DSU struct {
    parent []int
    rank []int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    rank := make([]int, n)

    for i:=0; i<n; i++ {
        parent[i] = i
    }
    return &DSU{parent, rank}
}

func(d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func(d *DSU) Union(x, y int) bool {
    rx, ry := d.Find(x), d.Find(y)
    if rx == ry {
        return false
    }

    if d.rank[rx] < d.rank[ry] {
        rx, ry = ry, rx
    }

    d.parent[ry] = rx
    if d.rank[rx] == d.rank[ry] {
        d.rank[rx]++
    }
    return true
}
