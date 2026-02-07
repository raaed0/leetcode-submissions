function furthestBuilding(heights: number[], bricks: number, ladders: number): number {
    let n = heights.length
    let maxHeap = new MaxPriorityQueue()

    for (let i = 1; i < n; i++) {
        let h = heights[i] - heights[i-1]
        if (h <= 0)
            continue

        bricks -= h
        maxHeap.enqueue(h)

        if (bricks < 0) {
            if (ladders > 0) {
                ladders--
                bricks += maxHeap.dequeue() as number
            } else {
                return i-1
            }

        }
    }

    return n-1
};
