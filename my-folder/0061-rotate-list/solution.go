/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
        return head
    }

    count, tail, next_tail := 0, head, head
    
    for tail.Next != nil {
        count++
        tail = tail.Next
    }

    count++
    k = k % count

    for i := 1; i < count - k; i++ {
        next_tail = next_tail.Next
    }

    tail.Next = head
    result_head := next_tail.Next
    next_tail.Next = nil

    return result_head
}
