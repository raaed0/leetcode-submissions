func kthSmallest(matrix [][]int, k int) int {
    minHeap := priorityqueue.NewWith(utils.IntComparator)
    n := len(matrix)
    var res interface{}
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            minHeap.Enqueue(matrix[i][j])
        }
    }

    for k >= 0 {
        if k == 0 {
           break
        }
        res, _ = minHeap.Dequeue()
        k--
    }
    return res.(int)
}
