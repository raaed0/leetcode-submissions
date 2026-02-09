function topKFrequent(nums: number[], k: number): number[] {
    let res = []
    let count = new Map<number, number>()

    for (const d of nums) {
        count.set(d, (count.get(d) ?? 0)+1)
    }

    let maxHeap = new MaxPriorityQueue<[number, number]>(c => c[0])
    for (const [num, freq] of count) {
        maxHeap.enqueue([freq, num])
    }

    while (k-- > 0) {
        res.push(maxHeap.dequeue()[1])
    }

    return res
};
