package trees

func levelOrder(root *TreeNode) [][]int {
	var levels [][]int
	traverse(root, 0, &levels)
	return levels
}

func traverse(node *TreeNode, level int, levels *[][]int) {
	if node == nil {
		return
	}

	if level == len(*levels) {
		*levels = append(*levels, []int{})
	}

	(*levels)[level] = append((*levels)[level], node.Val)
	level++
	traverse(node.Left, level, levels)
	traverse(node.Right, level, levels)
}
