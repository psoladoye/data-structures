package linkedlist

type dNode struct {
	data int
	prev *dNode
	next *dNode
}

func (d *dNode) GetData() int {
	return d.data
}

func newDoublyNode(data int) dNode {
	return dNode{data: data, next: nil, prev: nil}
}

type DoublyLinkedList struct {
	head *dNode
	tail *dNode
	size int
}

func (d *DoublyLinkedList) Head() Node {
	return d.head
}

func (d *DoublyLinkedList) Tail() Node {
	return d.tail
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

func (d *DoublyLinkedList) Prepend(data int) {

}

func (d *DoublyLinkedList) Append(data int) {

}

func (d *DoublyLinkedList) Pop() Node {
	return nil
}

func (d *DoublyLinkedList) Print() {

}
