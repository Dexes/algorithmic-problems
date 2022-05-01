package main

type MyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

func Constructor() MyLinkedList {
	head, tail := &Node{}, &Node{}
	head.Next, tail.Prev = tail, head

	return MyLinkedList{Head: head, Tail: tail, Length: 0}
}

func (x *MyLinkedList) Get(index int) int {
	if index >= x.Length {
		return -1
	}

	return x.getAtIndex(index).Val
}

func (x *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val, Prev: x.Head, Next: x.Head.Next}
	x.Head.Next, x.Head.Next.Prev = node, node
	x.Length++
}

func (x *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val, Prev: x.Tail.Prev, Next: x.Tail}
	x.Tail.Prev, x.Tail.Prev.Next = node, node
	x.Length++
}

func (x *MyLinkedList) AddAtIndex(index int, val int) {
	switch {
	case index == 0:
		x.AddAtHead(val)
	case index == x.Length:
		x.AddAtTail(val)
	case index < x.Length:
		current := x.getAtIndex(index)
		node := &Node{Val: val, Prev: current.Prev, Next: current}
		node.Prev.Next, node.Next.Prev = node, node
		x.Length++
	}
}

func (x *MyLinkedList) DeleteAtIndex(index int) {
	if index >= x.Length {
		return
	}

	current := x.getAtIndex(index)
	current.Prev.Next = current.Next
	current.Next.Prev = current.Prev
	x.Length--
}

func (x *MyLinkedList) getAtIndex(index int) *Node {
	switch {
	case index < x.Length/2:
		current, i := x.Head.Next, 0
		for ; i < index; current, i = current.Next, i+1 {
		}

		return current

	default:
		current, i := x.Tail.Prev, x.Length-1
		for ; i > index; current, i = current.Prev, i-1 {
		}

		return current
	}
}
