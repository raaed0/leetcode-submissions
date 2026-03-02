func countCompleteComponents(n int, edges [][]int) int {
    dsu := NewDSU(n)
    
    for _, e := range edges {
        dsu.Union(e[0], e[1])
    }

    edgeCount := map[int]int{}
    for _, e := range edges {
        root := dsu.Find(e[0])
        edgeCount[root]++
    }

    count := 0
    seen := map[int]bool{}

    for i:=0; i<n; i++ {
        root := dsu.Find(i)
        if seen[root] { continue }
        seen[root] = true

        k := dsu.GetSize(root)
        need := k*(k-1)/2
        if edgeCount[root] == need {
            count++
        }
    }

    return count
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
