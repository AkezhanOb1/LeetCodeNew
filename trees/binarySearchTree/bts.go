package binarysearchtree

//Node represents the each leaf of the tree
type Node struct {
	val   int
	right *Node
	left  *Node
}

//Bts - represents binary search trees
type Bts struct {
	root *Node
}

//GetRoot returns the root node of the tree
func (b Bts) GetRoot() *Node {
	return b.root
}

//Insert creates a node and places it in right
//position inside the tree
func (b Bts) Insert(value int) *Node {
	n := &Node{val: value}

	return n
}

func (b Bts) findCorrectPlace(value int) *Node {
	node := b.GetRoot()
	if node == nil {
		return b.root
	}

	//correctPlace := &Node{}
	for node != nil {
		if value > node.val {
			node = node.right
		} else {
			node = node.left
		}
	}
	return nil
}
