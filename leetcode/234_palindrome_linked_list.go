package leetcode

func isPalindromeList(head *ListNode) bool {
	rightPointer := reverseList(middleNode(head))
	leftPointer := head

	for rightPointer != nil {
		if rightPointer.Val != leftPointer.Val {
			return false
		}

		rightPointer = rightPointer.Next
		leftPointer = leftPointer.Next
	}

	return true
}
