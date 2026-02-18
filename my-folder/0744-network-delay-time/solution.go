type Edge struct {
    Target int
    Weight int
}

type Connection struct {
    Distance int
    Node int
}

func networkDelayTime(times [][]int, n int, k int) int {
    graph := make(map[int][]Edge)
    visited := make(map[int]struct{})
    res := 0

    for _, time := range times {
        origin := time[0]
        target := time[1]
        weight := time[2]

        graph[origin] = append(graph[origin], Edge {Target: target, Weight: weight})
    }
    minHeap := priorityqueue.NewWith(func(a, b interface{})int {
        conA := a.(Connection)
        conB := b.(Connection)
        return conA.Distance - conB.Distance
    })
    minHeap.Enqueue(Connection{Distance:0, Node: k})

    for minHeap.Size() > 0 {
        conn, _ := minHeap.Dequeue()
        current := conn.(Connection)
        w1, e1 := current.Distance, current.Node

        if _, exists := visited[e1]; exists {
            continue
        }
        visited[e1] = struct{}{}
        res = max(res, w1)

        for _, edge := range graph[e1] {
            target := edge.Target
            weight := edge.Weight

            if _, exists := visited[target]; !exists {
                minHeap.Enqueue(Connection{Distance: weight+w1, Node: target})
            }
        }
    }

    if len(visited) == n {
        return res
    }

    return -1
}
