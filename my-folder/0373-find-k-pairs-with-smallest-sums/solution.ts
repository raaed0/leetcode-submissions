function kSmallestPairs(nums1: number[], nums2: number[], k: number): number[][] {
    if (!nums1.length || !nums2.length || k === 0)
        return []

    
    let res: Array<[number, number]> = []
    let minHeap: MinPriorityQueue<[number, number, number]> = new MinPriorityQueue((pair) => pair[0])

    for (let j = 0; j < Math.min(nums2.length, k); j++) {
        minHeap.enqueue([nums1[0]+nums2[j], 0, j])
    }

    while (res.length < k && !minHeap.isEmpty()) {
        const [sum, i, j] = minHeap.dequeue()
        res.push([nums1[i], nums2[j]])

        if (i+1 < nums1.length) {
            minHeap.enqueue([nums1[i+1]+nums2[j], i+1, j])
        }
    }

    return res
};
