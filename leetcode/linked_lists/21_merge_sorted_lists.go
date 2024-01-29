package linked_lists

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	dummyNode := &ListNode{}
//
//	getNodeVal := func(node *ListNode) float64 {
//		if node == nil {
//			v := math.Inf(1)
//			return v
//		}
//
//		return float64(node.Val)
//	}
//
//	p1 := list1
//	p2 := list2
//	curr := dummyNode
//	for p1 != nil || p2 != nil {
//		if getNodeVal(p1) < getNodeVal(p2) {
//			curr.Next = p1
//			p1 = p1.Next
//		} else {
//			curr.Next = p2
//			p2 = p2.Next
//		}
//
//		curr = curr.Next
//	}
//
//	return dummyNode.Next
//}

// Решение без использования +infinity, но с более сложным условием в цикле.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}

	p1 := list1
	p2 := list2
	curr := dummyNode
	for p1 != nil || p2 != nil {
		if p1 != nil && (p2 == nil || p1.Val < p2.Val) {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}

		curr = curr.Next
	}

	return dummyNode.Next
}
