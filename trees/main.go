package main

import (
	"log"

	bst "./binarySearchTree"
)

func main() {
	tree := bst.BinarySearchTree{}
	tree.Constructor(5)
	tree.Insert(3)
	tree.Insert(6)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(7)

	log.Println(bst.FindTarget(tree.Root, 9))
}
