func makeConnected(n int, connections [][]int) int {
    dsu := NewDSU(n)
    nc := n-1
    extra := 0

    for _, e := range connections {
        rx, ry := dsu.Find(e[0]), dsu.Find(e[1])
        if rx == ry {
            extra++
        } else {
            dsu.Union(e[0], e[1])
            nc--
        }
    }

    if nc <= extra {
        return nc
    } else {
        return -1
    }
}

type DSU struct {
    parent []int
    rank   []int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    rank := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = i
    }
    return &DSU{parent, rank}
}

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) {
    rx, ry := d.Find(x), d.Find(y)
    if rx == ry { return }
    if d.rank[rx] < d.rank[ry] {
        rx, ry = ry, rx
    }
    d.parent[ry] = rx
    if d.rank[rx] == d.rank[ry] {
        d.rank[rx]++
    }
}
