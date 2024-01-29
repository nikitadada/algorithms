package linked_lists

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	dummyNode := &ListNode{}
//	additionalNumber := 0
//	currVal := 0
//
//	currNode := dummyNode
//	for l1 != nil || l2 != nil {
//		if l1 != nil {
//			currVal += l1.Val
//			l1 = l1.Next
//		}
//		if l2 != nil {
//			currVal += l2.Val
//			l2 = l2.Next
//		}
//
//		currVal += additionalNumber
//		if currVal > 9 {
//			currVal %= 10
//			additionalNumber = 1
//		} else {
//			additionalNumber = 0
//		}
//
//		node := &ListNode{
//			Val:  currVal,
//			Next: nil,
//		}
//		currNode.Next = node
//		currNode = currNode.Next
//		currVal = 0
//	}
//
//	if additionalNumber == 1 {
//		node := &ListNode{
//			Val:  1,
//			Next: nil,
//		}
//		currNode.Next = node
//	}
//
//	return dummyNode.Next
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	currVal := 0

	for currNode := dummyNode; l1 != nil || l2 != nil || currVal != 0; currNode = currNode.Next {
		if l1 != nil {
			currVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currVal += l2.Val
			l2 = l2.Next
		}

		node := &ListNode{Val: currVal % 10}
		currNode.Next = node
		currVal /= 10
	}

	return dummyNode.Next
}
