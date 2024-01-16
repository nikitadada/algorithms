package trees

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root.Left, root.Right)
}

func checkSymmetric(l *TreeNode, r *TreeNode) bool {
	if l == nil || r == nil {
		return l == r
	}

	if l.Val != r.Val {
		return false
	}

	return checkSymmetric(l.Left, r.Right) && checkSymmetric(l.Right, r.Left)
}
