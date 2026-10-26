/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    i:= 0
    j:= 0
    dummy := head
    current:= head
    for dummy != nil{
        i++
        dummy= dummy.Next
    }
    if i == n {
        return head.Next
    }
    for current.Next != nil {
       
       
        if j == i-n -1{
           current.Next = current.Next.Next
           return head
        }

        j++
        current= current.Next
    }
    return head
}