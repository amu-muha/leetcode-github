/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    slow := head
    fast:= head
    for fast != nil {   
        if fast.Next == nil {
            return slow
        }
    fast = fast.Next.Next
    slow = slow.Next
    }
    return slow
}