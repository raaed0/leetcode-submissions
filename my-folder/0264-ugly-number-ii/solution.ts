function nthUglyNumber(n: number): number {
    let minHeap = new MinPriorityQueue<number>()
    minHeap.enqueue(1)

    let factors: number[] = [2, 3, 5]
    let set = new Set<number>()

    for (let i = 0; i < n; i++) {
        let num = minHeap.dequeue()
        if (i === n-1)
            return num

        for (let f of factors) {
            let t = f * num
            if (!set.has(t)) {
                set.add(t)
                minHeap.enqueue(t)
            }
        }      
    }

};


