/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    groupPrev := dummy

    for {
        // Find the kth node in the current group
        kth := groupPrev
        for i := 0; i < k; i++ {
            kth = kth.Next
            if kth == nil {
                return dummy.Next
            }
        }

        groupNext := kth.Next

        // Reverse the group
        prev := groupNext
        curr := groupPrev.Next

        for curr != groupNext {
            next := curr.Next
            curr.Next = prev
            prev = curr
            curr = next
        }

        // Connect the reversed group
        temp := groupPrev.Next
        groupPrev.Next = kth
        groupPrev = temp
    }
}
