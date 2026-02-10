function frequencySort(s: string): string {
    let count: Map<string, number> = new Map()
    let maxHeap: MaxPriorityQueue<[number, string]> = new MaxPriorityQueue(v => v[0])
    let res: string = ""

    for (let c of s) {
        count.set(c, (count.get(c) ?? 0)+1)
    }

    for (const [ch, freq] of count) {
        maxHeap.enqueue([freq, ch])
    }

    while (maxHeap.size() > 0) {
        let [freq, ch] = maxHeap.dequeue()
        res += ch.repeat(freq)
    }

    return res
};
