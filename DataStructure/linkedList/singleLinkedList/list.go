package singlelinkedlist

import "fmt"

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

//Push is the method of of the list that puts a new data to the Linked List
//and returns the length of the entire list
func (l *List) Push(value int) int {
	newNode := &Node{
		Val: value,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Len++
		return l.Len
	}

	node := l.Head
	for node.Next != nil {
		node = node.Next
	}
	node.Next = newNode
	l.Tail = newNode
	l.Len++
	return l.Len
}

//Traverse is the list method that shows all the node's value
//that stored inside the list
func (l *List) Traverse() {
	node := l.Head
	for node != nil {
		fmt.Println(node)
		node = node.Next
	}
}

//Pop is the list method that removes the very first node from the list
//and returns it's value, the the list does not contain any value it
//return 0 and false otherwise the return values are node val and true
func (l *List) Pop() (bool, int) {
	head := l.Head
	if head == nil {
		return false, 0
	}

	l.Head = head.Next
	l.Len--
	return true, head.Val
}

//Shift is the list method that puts a new value to the end of the list
//and returns the length of the list
func (l *List) Shift(value int) int {
	newNode := &Node{
		Val: value,
	}

	l.Tail.Next = newNode
	l.Tail = newNode
	l.Len++
	return l.Len
}

//Unshift is the method that removes the very last node from the list
//returns it's value and puts the last but one node as list's tail
func (l *List) Unshift() (bool, int) {
	tail := l.Tail
	if tail == nil {
		return false, 0
	}
	node := l.Head
	for node.Next != tail {
		node = node.Next
	}
	l.Tail = node
	val := node.Val
	node.Next = nil
	l.Len--
	return true, val

}

//Reverse is the list method that reverses the entire list
//very last node becomes very first
func (l *List) Reverse() {
	if l.Len == 0 {
		return
	}
	//6-5-3-4-2-1
	node := l.Head
	l.Head, l.Tail = l.Tail, l.Head

	var next *Node
	var prev *Node
	for i := 0; i < l.Len; i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
}

//MiddleNode returns the middle of the list
func (l *List) MiddleNode() *Node {
	counter := 0
	head := l.Head
	for head != nil {
		head = head.Next
		counter++
	}

	node := l.Head
	for i := 0; i < counter/2; i++ {
		node = node.Next
	}
	return node
}
