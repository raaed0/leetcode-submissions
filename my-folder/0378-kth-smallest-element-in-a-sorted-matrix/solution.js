/**
 * @param {number[][]} matrix
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function(matrix, k) {
    let minHeap = new MinPriorityQueue(v => v[0])
    let n = matrix.length
    let res = 0

    for (let i = 0; i < n; i++) {
        minHeap.enqueue([matrix[i][0], i, 0])
    }

    while (k-- > 0) {
        const [v, r, c] = minHeap.dequeue()
        res = v

        if (c+1 < n) {
            minHeap.enqueue([matrix[r][c+1], r, c+1])
        }
    }

    return res
};
