package trees

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum = root.Left.Val
	}

	return sum + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
