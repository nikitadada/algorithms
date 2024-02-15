package trees

type Queue struct {
	items []*TreeNode
}

func (q *Queue) push(node *TreeNode) {
	q.items = append(q.items, node)
}
func (q *Queue) pop() *TreeNode {
	deleted := q.items[0]
	q.items = q.items[1:]

	return deleted
}
func (q *Queue) len() int {
	return len(q.items)
}

// Решение с помощью BFS
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := &Queue{}
	queue.push(root)

	leftToRight := true
	for queue.len() > 0 {
		levelSize := queue.len()
		currentLevel := make([]int, 0)
		for i := 0; i < levelSize; i++ {
			currentNode := queue.pop()

			if leftToRight {
				currentLevel = append(currentLevel, currentNode.Val)
			} else {
				currentLevel = append([]int{currentNode.Val}, currentLevel...)
			}

			if currentNode.Left != nil {
				queue.push(currentNode.Left)
			}
			if currentNode.Right != nil {
				queue.push(currentNode.Right)
			}
		}

		res = append(res, currentLevel)
		leftToRight = !leftToRight
	}

	return res
}
