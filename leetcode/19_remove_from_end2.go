package leetcode

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Next: head,
	}
	fast := dummyNode
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}

	slow := dummyNode
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummyNode.Next
}
