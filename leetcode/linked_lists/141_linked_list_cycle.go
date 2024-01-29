package linked_lists

func hasCycle(head *ListNode) bool {
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
