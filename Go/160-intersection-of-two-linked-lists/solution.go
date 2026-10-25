/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    check := make(map[*ListNode]bool)
    for headA != nil {
        check[headA] = true
        headA = headA.Next
    }

    for headB != nil {
        if check[headB]{
            return headB
        }
        headB = headB.Next
    }
    return nil
}