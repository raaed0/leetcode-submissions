func accountsMerge(accounts [][]string) [][]string {
    dsu := NewDSU(len(accounts))

    emailToId := make(map[string]int)

    for i, acc := range accounts {
        for _, e := range acc[1:] {
            if id, ok := emailToId[e]; ok {
                dsu.Union(i, id)
            } else {
                emailToId[e] = i
            }
        }
    }

    emailGroup := make(map[int][]string)

    for e, id := range emailToId {
        root := dsu.Find(id)

        emailGroup[root] = append(emailGroup[root], e)
    }

    res := [][]string{}

    for id, emails := range emailGroup {
        name := accounts[id][0]
        sort.Strings(emails)
        emails = append([]string{name}, emails...)
        res = append(res, emails)
    }

    return res
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
