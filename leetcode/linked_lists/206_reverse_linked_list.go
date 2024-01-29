package linked_lists

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	cur := head

	for cur != nil {
		tmp := cur

		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}
	head = prev

	return head
}
