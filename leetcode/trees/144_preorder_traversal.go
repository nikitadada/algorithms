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

//func preorderTraversal(root *TreeNode) []int {
//	var res []int
//	traverse(root, &res)
//	return res
//}
//
//func traverse (node *TreeNode, res *[]int) {
//	if node == nil {
//		return
//	}
//
//	*res = append(*res, node.Val)
//	traverse(node.Left, res)
//	traverse(node.Right, res)
//}
