/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    fast := head
    slow:= head
    for fast != nil && fast.Next != nil{
      fast =  fast.Next.Next
      slow =  slow.Next
    }
    
    var prev *ListNode = nil
    current := slow.Next
    slow.Next = nil

    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }

    first := head
    second := prev

    for second != nil{
        firstNext := first.Next
        secondNext := second.Next

        first.Next = second
        second.Next = firstNext

        first = firstNext
        second = secondNext
    }

}