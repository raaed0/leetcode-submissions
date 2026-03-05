func equationsPossible(equations []string) bool {
	dsu := NewDSU(26)

	for _, str := range equations {
		if str[1] == '=' {
			dsu.Union(int(str[0]-'a'), int(str[3]-'a'))
		}
	}

	for _, str := range equations {
		if str[1] == '!' && (dsu.Find(int(str[0]-'a')) == dsu.Find(int(str[3]-'a'))) {
			return false
		}
	}

	return true
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

func (d *DSU) Union(x, y int) bool {
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
