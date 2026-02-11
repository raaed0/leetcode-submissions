func topKFrequent(words []string, k int) []string {
    type Entry struct {
        str string
        freq int
    }

    count := make(map[string]int)
    maxheap := priorityqueue.NewWith(func(a, b interface{}) int {
        ea, eb := a.(Entry), b.(Entry)
        if ea.freq == eb.freq {
            if ea.str < eb.str {
                return -1
            } else {
                return 1
            }
        }
        return eb.freq - ea.freq
    })

    for _, s := range words {
        count[s]++
    }

    for str, freq := range count {
        maxheap.Enqueue(Entry{str, freq})
    }

    res := make([]string, 0, k)
    for k > 0 {
        val, _ := maxheap.Dequeue()
        res = append(res, val.(Entry).str)
        k--
    }

    return res
}
