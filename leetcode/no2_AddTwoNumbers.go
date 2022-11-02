func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resNode := ListNode{0, nil}
	prevNode := &resNode
	tmpNode := ListNode{0, nil}
	var currNode *ListNode
	carry := 0
	for {        
			sum := l1.Val + l2.Val + carry
			curr := sum % 10
			carry = sum / 10
			
			currNode = &ListNode{curr, nil}
			prevNode.Next = currNode
			prevNode = currNode
			
			flag := 0
			if l1.Next == nil {
					l1 = &tmpNode
					flag++
			} else {
					l1 = l1.Next
			}
			if l2.Next == nil {
					l2 = &tmpNode
					flag++
			} else {
					l2 = l2.Next
			}
			
			if flag == 2 && carry == 0 {
					break
			}
	}
	return resNode.Next  
}