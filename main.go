package main

import linkedList "github.com/psoladoye/datastructures/linkedlist"

func main() {
	list := linkedList.LinkedList(&linkedList.SinglyLinkedList{})
	list.Prepend(linkedList.NewNode(10))
	list.Prepend(linkedList.NewNode(24))
	list.Prepend(linkedList.NewNode(1))
	list.Prepend(linkedList.NewNode(35))
	list.Print()
}
