function findLeastNumOfUniqueInts(arr: number[], k: number): number {
    let count = new Map<number, number>()

    for (let d of arr) {
        count.set(d, (count.get(d) ?? 0)+1)
    }

    let minHeap = new MinPriorityQueue()
    for (const [_, value] of count) {
        minHeap.push(value)
    }

    while (k > 0) {
        let t = minHeap.pop() as number
        if (k >= t) {
            k = k - t
            continue
        }
        t = t - k as number
        minHeap.push(t)
        k = 0
    }

    return minHeap.size()
};
