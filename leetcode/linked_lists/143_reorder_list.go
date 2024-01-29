package linked_lists

func reorderList(head *ListNode) {
	right := reverseList(middleNode(head))

	left := head
	for left != nil && right != left.Next {
		leftNext := left.Next
		left.Next = right
		left = right
		right = leftNext
	}
}
