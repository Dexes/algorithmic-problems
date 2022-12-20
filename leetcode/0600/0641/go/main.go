package main

const (
	front = iota
	back
)

type MyCircularDeque struct {
	head, tail       *Node
	length, capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{capacity: k}
}

func (x *MyCircularDeque) InsertFront(value int) bool {
	return x.insertNode(value, front)
}

func (x *MyCircularDeque) InsertLast(value int) bool {
	return x.insertNode(value, back)
}

func (x *MyCircularDeque) DeleteFront() bool {
	return x.deleteNode(front)
}

func (x *MyCircularDeque) DeleteLast() bool {
	return x.deleteNode(back)
}

func (x *MyCircularDeque) GetFront() int {
	if x.IsEmpty() {
		return -1
	}

	return x.head.Value
}

func (x *MyCircularDeque) GetRear() int {
	if x.IsEmpty() {
		return -1
	}

	return x.tail.Value
}

func (x *MyCircularDeque) IsEmpty() bool {
	return x.head == nil
}

func (x *MyCircularDeque) IsFull() bool {
	return x.length == x.capacity
}

func (x *MyCircularDeque) insertNode(value, direction int) bool {
	if x.IsFull() {
		return false
	}

	x.length++
	node := &Node{Value: value}

	if x.tail == nil {
		x.head, x.tail = node, node
		node.Prev, node.Next = node, node
		return true
	}

	node.Prev, node.Next = x.tail, x.head
	x.tail.Next, x.head.Prev = node, node

	if direction == front {
		x.head = node
	} else {
		x.tail = node
	}

	return true
}

func (x *MyCircularDeque) deleteNode(direction int) bool {
	if x.IsEmpty() {
		return false
	}

	x.length--

	if x.head == x.tail {
		x.head, x.tail = nil, nil
		return true
	}

	if direction == front {
		x.head = x.head.Next
	} else {
		x.tail = x.tail.Prev
	}

	x.tail.Next, x.head.Prev = x.head, x.tail
	return true
}

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
}
