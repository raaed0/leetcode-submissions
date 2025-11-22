/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    visited := make(map[*ListNode]bool)

    for curr := headA; curr != nil; curr = curr.Next {
        visited[curr] = true
    }

    for curr := headB; curr != nil; curr = curr.Next {
        if visited[curr] {
            return curr
        }
    }

    return nil
}
