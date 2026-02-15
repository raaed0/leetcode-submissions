func findClosestElements(arr []int, k int, x int) []int {
    type Element struct {
        d int
        distance int
    }

    maxHeap := priorityqueue.NewWith(func(a, b interface{}) int {
        da, db := a.(Element), b.(Element)
        if da.distance == db.distance {
            return db.d - da.d
        }
        return db.distance - da.distance
    })

    for _, d := range arr {
        maxHeap.Enqueue(Element{d: d, distance: abs(d-x)})
        for maxHeap.Size() > k {
            maxHeap.Dequeue()
        }
    }

    res := make([]int, 0, k)
    for k > 0 {
        el, _ := maxHeap.Dequeue()
        res = append([]int{el.(Element).d}, res...)
        k--
    }
    sort.Ints(res)
    return res
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
