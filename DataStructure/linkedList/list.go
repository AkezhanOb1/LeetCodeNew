package linkedlist

//Node is the type that represents the singe peace of some data
//in the linked list, it also cointains the reference to the next node
//in the list
type Node struct {
	Val  int
	Next *Node
}

//List represent the entire linked list structure
//it has fileds length - the lengh of the list tail the
//reference to the very last node int the list and head
//the reference to the very fist node in the linked list
type List struct {
	Head *Node
	Tail *Node
	Len  int
}
