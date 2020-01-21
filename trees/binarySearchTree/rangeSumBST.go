package binarysearchtree

//RangeSumBST is
func RangeSumBST(root *Node, L int, R int) int {
	counter := 0
	traverseHelper(root, &counter, L, R)
	return counter
}

func traverseHelper(node *Node, counter *int, L int, R int) {
	if node.Val >= L && node.Val <= R {
		*counter += node.Val
	}

	if node.Left != nil {
		traverseHelper(node.Left, counter, L, R)
	}

	if node.Right != nil {
		traverseHelper(node.Right, counter, L, R)
	}

}
