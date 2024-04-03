package linkedlist

import (
	"fmt"
	"log"
)

type sNode struct {
	data int
	next *sNode
}

func (n *sNode) GetData() int {
	return n.data
}

type SinglyLinkedList struct {
	head *sNode
	tail *sNode
	size int
}

func newSinglyNode(data int) sNode {
	return sNode{data: data, next: nil}
}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Head() Node {
	return s.head
}

func (s *SinglyLinkedList) Tail() Node {
	return s.tail
}

func (s *SinglyLinkedList) Prepend(data int) {
	node := newSinglyNode(data)
	node.next = s.head
	s.head = &node

	if s.size == 0 {
		s.tail = &node
	}

	s.size++
}

func (s *SinglyLinkedList) Append(data int) {
	node := newSinglyNode(data)
	if s.size != 0 {
		lastNode := s.tail
		lastNode.next = &node
	} else {
		s.head = &node
	}

	s.tail = &node
	s.size++
}

func (s *SinglyLinkedList) Pop() Node {
	node := *s.head
	s.head = s.head.next
	s.size--

	if s.size == 0 {
		s.tail = nil
	}

	return &node
}

func (s *SinglyLinkedList) Print() {
	ptr := s.head
	output := "(head)"
	for ptr != nil {
		n := fmt.Sprintf("-->(%d|%v)", ptr.data, ptr.next)
		output += n
		ptr = ptr.next
	}
	log.Printf("Size => %d\n", s.size)
	log.Println(output)
}
