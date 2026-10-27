/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    dummy := &ListNode{Next: head}

    length := 0
    current := head

    for current != nil {
        length++
        current = current.Next
    }

    k = k % length

    for i := 0; i < k; i++ {
        current = dummy

        for current.Next.Next != nil {
            current = current.Next
        }

        node := current.Next
        current.Next = nil
        node.Next = dummy.Next
        dummy.Next = node
    }

    return dummy.Next
}