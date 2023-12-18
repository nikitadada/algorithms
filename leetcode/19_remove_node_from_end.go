package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	curr := dummyNode
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}

	curr = dummyNode
	for i := 0; i < count-n-1; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return dummyNode.Next
}
