/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function getIntersectionNode(headA: ListNode | null, headB: ListNode | null): ListNode | null {
    if (headA == null || headB == null) {
        return null
    }

    let pA = headA
    let pB = headB

    while (pA != pB) {
        if (pA == null) {
            pA = headB
        } else {
            pA = pA.next
        }

        if (pB == null) {
            pB = headA
        } else {
            pB = pB.next
        }
    }

    return pA
};
