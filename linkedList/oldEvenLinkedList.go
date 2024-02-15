// maintain 2 pointer, even and odd pointer
// resolve this in O(n) time complexity
//maintain 2 grp, one for odd number and one for even
//use total 4 pointer, 2 for maintaining head and 2 for iteration 
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    oddHead, evenHead := head, head.Next
    odd, even := oddHead, evenHead

    for even != nil && even.Next != nil{
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead

    return oddHead

}
