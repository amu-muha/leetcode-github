/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    previous := dummy

    for previous.Next != nil && previous.Next.Next != nil {
        first := previous.Next
        second := first.Next
        previous.Next = second
        first.Next = second.Next
        second.Next = first

        previous = first
    }

    return dummy.Next
}