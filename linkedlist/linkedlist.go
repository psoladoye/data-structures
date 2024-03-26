package linkedlist

import (
	"fmt"
	"log"
)

type Node struct {
	data int
	next *Node
}

func NewNode(data int) Node {
	return Node{data: data, next: nil}
}

type LinkedList interface {
	Prepend(Node)
	Pop()
	Print()
}

type SinglyLinkedList struct {
	head *Node
}

func (s *SinglyLinkedList) Prepend(node Node) {
	node.next = s.head
	s.head = &node
}

func (s *SinglyLinkedList) Pop() {
	s.head = s.head.next
}

func (s *SinglyLinkedList) Print() {
	ptr := s.head
	output := "(head)"
	for ptr != nil {
		n := fmt.Sprintf("-->(%d|%v)", ptr.data, ptr.next)
		output += n
		ptr = ptr.next
	}
	log.Println(output)
}
