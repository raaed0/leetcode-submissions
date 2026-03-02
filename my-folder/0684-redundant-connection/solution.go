func findRedundantConnection(edges [][]int) []int {
    n := len(edges)
    dsu := NewDSU(n+1)

    for _, e := range edges {
        if !dsu.Union(e[0], e[1]) {
            return e
        }
    }

    return nil
}

type DSU struct {
    parent []int
    size []int
}

func NewDSU(n int) *DSU {
    parent := make([]int, n)
    size := make([]int, n)

    for i := 0; i<n; i++ {
        parent[i] = i
        size[i] = 1
    }

    return &DSU{parent, size}
}

func (d *DSU) Find(x int) int {
    if d.parent[x] != x {
        d.parent[x] = d.Find(d.parent[x])
    }
    return d.parent[x]
}

func (d *DSU) Union(x, y int) bool {
    rx, ry := d.Find(x), d.Find(y)
    if rx == ry { return false}

    if d.size[rx] < d.size[ry] { rx, ry = ry, rx}
    d.parent[ry] = rx
    d.size[rx] += d.size[ry]
    return true
}

func (d *DSU) GetSize(x int) int {
    return d.size[d.Find(x)]
}
