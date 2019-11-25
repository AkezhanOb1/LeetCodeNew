package main

import (
	l "./singleLinkedList"
)

func main() {
	list := l.List{}
	list.Push(6)
	list.Push(5)
	list.Push(4)
	list.Push(3)
	list.Push(2)
	list.Push(1)
	list2 := l.List{}
	list2.Push(7)
	list2.Push(8)
	list2.Push(9)
	list2.Push(10)
	list2.Push(11)
	list2.Push(12)
	merges := l.Merge(list, list2)
	merges.Traverse()
}
