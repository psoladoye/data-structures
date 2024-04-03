package linkedlist

type Node interface {
	GetData() int
}

type LinkedList interface {
	Head() Node
	Tail() Node
	Prepend(int)
	Append(int)
	Pop() Node
	Print()
	Size() int
}
