/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    
    p1, p2 := head, head
    for p2.Next != nil && p2.Next.Next != nil { 
        p1 = p1.Next
        p2 = p2.Next.Next
    }    
    head2 := p1.Next
    p1.Next = nil
     
    front, mid, end := head2, head2.Next, head2.Next
    front.Next = nil
    for mid != nil {
        end = mid.Next
        mid.Next = front
        front = mid
        mid = end
    }
    head2 = front

    
    front, mid, end = head, head2, head.Next 
    for mid != nil {
        front.Next = mid
        front = mid
        mid = end
        end = front.Next
    }
}
