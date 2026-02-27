type Point struct {
    Denominator string
    Value float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    adj := make(map[string][]Point)

    for i, e := range equations {
        adj[e[0]] = append(adj[e[0]], Point{e[1], values[i]})
        adj[e[1]] = append(adj[e[1]], Point{e[0], 1 / values[i]})
    }

    var bfs func(src, target string) float64
    bfs = func(src, target string) float64 {
        if _, ok := adj[src]; !ok {
            return -1.0
        }
        if _, ok := adj[target]; !ok {
            return -1.0
        }

        q := [][2]interface{}{{src, 1.0}}
        visited := make(map[string]bool)
        visited[src] = true

        for len(q) > 0 {
            node := q[0][0].(string)
            weight := q[0][1].(float64)
            q = q[1:]

            if node == target {
                return weight
            }

            for _, ngh := range adj[node] {
                nghNode := ngh.Denominator
                nghWeight := ngh.Value

                if !visited[nghNode] {
                    visited[nghNode] = true
                    q = append(q, [2]interface{}{nghNode, weight * nghWeight})
                }
            }
        }

        return -1.0
    }

    res := make([]float64, len(queries))
    for i, q := range queries {
        res[i] = bfs(q[0], q[1])
    }

    return res
}

