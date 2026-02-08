function findKthLargest(nums: number[], k: number): number {
    if (k === 0)
        return 0

    let maxHeap = new MaxPriorityQueue<number>()

    for (const num of nums) {
        maxHeap.enqueue(num)
    }

    while (k-- >= 0) {
        let cur = maxHeap.dequeue()
        if (k === 0) {
            return cur
        }
    }
};
