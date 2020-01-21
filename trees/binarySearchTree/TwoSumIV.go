package binarysearchtree

//FindTarget is
func FindTarget(root *Node, k int) bool {
	m := make(map[int]bool)

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := m[node.Val]; ok {
			return true
		}

		m[k-node.Val] = true

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}
