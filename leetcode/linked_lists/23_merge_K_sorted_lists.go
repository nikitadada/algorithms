package linked_lists

func mergeKLists(lists []*ListNode) *ListNode {
	for len(lists) > 1 {
		list1 := lists[0]
		list2 := lists[1]
		lists := lists[2:]

		mergedTwoLists := mergeTwoLists2(list1, list2)
		lists = append(lists, mergedTwoLists)
	}

	return lists[0]
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
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
