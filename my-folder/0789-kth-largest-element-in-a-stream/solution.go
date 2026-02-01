type KthLargest struct {
    k int
    heap []int
}


func Constructor(k int, nums []int) KthLargest {
    return KthLargest{k: k, heap: nums}
}

func (this *KthLargest) Add(val int) int  {
    this.heap = append(this.heap, val)
    sort.Ints(this.heap)
    return this.heap[len(this.heap)-this.k]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
