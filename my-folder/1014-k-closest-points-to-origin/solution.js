/**
 * @param {number[][]} points
 * @param {number} k
 * @return {number[][]}
 */
var kClosest = function(points, k) {
    const minHeap = new MinPriorityQueue((point) => point[0]);

    for (let [x,y] of points) {
        let dist = x ** 2 + y ** 2;
        minHeap.enqueue([dist, x, y])
    }

    const result = [];
    for (let i = 0; i < k; i++) {
        const [_, x, y] = minHeap.dequeue();
        result.push([x, y])
    }

    return result;
};
