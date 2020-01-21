package binarysearchtree

//KthSmallest is
func KthSmallest(root *Node, k int) int {
	res := 0
	helper(root, &k, &res)
	return res
}

func helper(root *Node, k *int, res *int) {
	if *res != 0 {
		return
	}
	if root.Left != nil {
		helper(root.Left, k, res)
	}
	*k--
	if *k == 0 {
		*res = root.Val
		return
	}
	if root.Right != nil {
		helper(root.Right, k, res)
	}
}
