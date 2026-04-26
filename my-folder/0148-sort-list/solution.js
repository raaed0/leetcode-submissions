/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var sortList = function(head) {
    if (!head || !head.next) {
        return head
    }

    let length = 0
    let curr = head

    while(curr) {
        length++
        curr = curr.next
    }

    let dummy = new ListNode(0)
    dummy.next = head
    let step = 1

    while (step < length) {
        let prev = dummy,
            curr = dummy.next;
        
        while (curr) {
            let left = curr
            let right = split(left, step)
            curr = split(right, step)
            let merged = merge(left, right)
            prev.next = merged
            while (prev.next) {
                prev = prev.next
            }
        }

        step *= 2
    }

    return dummy.next
};

function split(head, step) {
    if (!head) return null
    for (let i = 0; i < step-1 && head.next; i++) {
        head = head.next
    }
    let nextPart = head.next
    head.next = null
    return nextPart
}

function merge(list1, list2) {
    let dummy = new ListNode(0)
    let tail = dummy

    while (list1 && list2) {
        if (list1.val < list2.val) {
            tail.next = list1
            list1 = list1.next
        } else {
            tail.next = list2
            list2 = list2.next
        }
        tail = tail.next
    }

    tail.next = list1 || list2
    return dummy.next
}
