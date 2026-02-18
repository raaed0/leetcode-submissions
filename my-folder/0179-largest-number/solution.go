func largestNumber(nums []int) string {
    res := ""

    maxHeap := priorityqueue.NewWith(func(a, b interface{})int {
        r1 := a.(string) + b.(string)
        r2 := b.(string) + a.(string)

        if r1 > r2 {
            return -1
        } else {
            return 1
        }
    })

    for _, n := range nums {
        maxHeap.Enqueue(strconv.Itoa(n))
    }

    for maxHeap.Size() > 0 {
        v, _ := maxHeap.Dequeue()
        res += v.(string)
    }

    if res[0] == '0' {
        return "0"
    }
    return res
}
