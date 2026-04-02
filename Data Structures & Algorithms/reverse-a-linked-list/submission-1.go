/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
  if head == nil|| head.Next == nil{
    return head
  }  

  curr := reverseList(head.Next)
  head.Next.Next = head
  head.Next = nil

  return curr
}

// func reverseList(head *ListNode) *ListNode {
//     var prev *ListNode
//     curr := head
//     for curr != nil {
//         nxt := curr.Next
//         curr.Next = prev
//         prev = curr
//         curr = nxt
//     }
//     return prev
// }
