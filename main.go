package main

import . "github.com/psoladoye/datastructures/linkedlist"

func main() {
	linkedList := LinkedList(&SinglyLinkedList{})
	linkedList.Prepend(NewNode(10))
	linkedList.Prepend(NewNode(24))
	linkedList.Prepend(NewNode(1))
	linkedList.Prepend(NewNode(35))
	linkedList.Print()
}
