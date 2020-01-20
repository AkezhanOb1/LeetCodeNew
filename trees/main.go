package main

import (
	"log"

	bts "./binarySearchTree"
)

func main() {
	tree := bts.BinarySearchTree{}
	tree.Constructor(10)

	tree.Insert(6)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(8)
	tree.Insert(20)

	log.Println(tree.PreOrderDFS())
}
