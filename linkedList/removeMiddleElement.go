// remember head is pointer so new listNode is created with head is a reference to it
// changes in newList node will effect the orignal head
// find length of linkedList then delete it using the below 

func deleteMiddle(head *ListNode) *ListNode {
   var length int
     newNode := head
    currentNode := head
  
   for newNode != nil {
      newNode = newNode.Next
       length++   
    }
    
    if length == 1{
        return nil
    }


    middleNode := length/2

    for i:=0;i<middleNode-1;i++{
         currentNode = currentNode.Next
    }
   
    currentNode.Next = currentNode.Next.Next
    return head

}

func main() {
	// Example usage
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	newHead := deleteMiddle(head)
	// Use the newHead as needed
}
