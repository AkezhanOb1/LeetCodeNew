package binarysearchtree

// Node - represents a single node in the binary search tree
// it has 3 properties: Val - actual value that is stores,
// link to the left node which should be less than current node,
// link to the right node which should be more than current node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// BinarySearchTree - represents a BTS and contains only one properties
// link to the root(very first) node
type BinarySearchTree struct {
	Root *Node
}

//Constructor - creates a new bst and sets a root node
//with the provided value and returns that node
func (b *BinarySearchTree) Constructor(val int) *Node {
	if ok := b.CheckRoot(); ok {
		return nil
	}

	b.Root = &Node{
		Val: val,
	}

	return b.Root
}

//CheckRoot - is method for determining the existence of
//the root node, if root is set then returns true
//otherwise it returns false
func (b BinarySearchTree) CheckRoot() bool {
	if b.Root == nil {
		return false
	}
	return true
}

//Insert - is a method that creates a node with the
//given value and inserts it into the tree
func (b *BinarySearchTree) Insert(val int) *Node {
	if ok := b.CheckRoot(); !ok {
		return b.Constructor(val)
	}

	newNode := &Node{
		Val: val,
	}

	node := b.Root

	for node != nil {
		if node.Val > val {
			if node.Left == nil {
				node.Left = newNode
				break
			}
			node = node.Left
		} else if node.Val < val {
			if node.Right == nil {
				node.Right = newNode
				break
			}
			node = node.Right
		} else {
			break
		}
	}

	return node
}

//Find is the method that searches a tree for specific
//value, if it exists returns a node, otherwise returns
//a nil
func (b BinarySearchTree) Find(val int) *Node {

	if ok := b.CheckRoot(); !ok {
		return nil
	}

	node := b.Root
	for node != nil {
		if val == node.Val {
			return node
		}

		if val > node.Val {
			if node.Right == nil {
				return nil
			}
			node = node.Right
		} else if val < node.Val {
			if node.Left == nil {
				return nil
			}
			node = node.Left
		}
	}

	return nil
}

//BFS is the method for traversing the bst
//from top to the bottom from left to right, bredth
func (b BinarySearchTree) BFS() []int {
	if ok := b.CheckRoot(); !ok {
		return nil
	}

	var queue []*Node
	var visited []int

	queue = append(queue, b.Root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		visited = append(visited, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return visited
}

//PreOrderDFS is the method for traversing the bst
//from top to bottom, left node first
func (b BinarySearchTree) PreOrderDFS() []int {
	if ok := b.CheckRoot(); !ok {
		return nil
	}

	var result []int
	node := b.Root
	preOrderDFSHelper(node, &result)
	return result
}

func preOrderDFSHelper(node *Node, res *[]int) {
	*res = append(*res, node.Val)
	if node.Left != nil {
		preOrderDFSHelper(node.Left, res)
	}
	if node.Right != nil {
		preOrderDFSHelper(node.Right, res)
	}
}
