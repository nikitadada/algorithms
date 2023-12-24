package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	Res []int
}

func preorderTraversal(root *TreeNode) []int {
	r := Res{}
	r.traverse(root)
	return r.Res
}

func (r *Res) traverse(node *TreeNode) {
	if node == nil {
		return
	}

	r.Res = append(r.Res, node.Val)
	r.traverse(node.Left)
	r.traverse(node.Right)
}
