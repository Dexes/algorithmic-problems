package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type Queue []*Node

func (q Queue) Push(nodes ...*Node) Queue {
	for _, node := range nodes {
		if node != nil {
			q = append(q, node)
		}
	}

	return q
}

func (q Queue) Pop() (*Node, Queue) {
	return q[0], q[1:]
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func makeQueue(root *Node, capacity int) Queue {
	result := make(Queue, 0, capacity)
	return result.Push(root)
}

func connect(root *Node) *Node {
	var (
		level, levelLength = 1, 1
		node, right        *Node
		queue              = makeQueue(root, 1<<12-1)
	)

	for !queue.IsEmpty() {
		node, queue = queue.Pop()
		queue = queue.Push(node.Right, node.Left)
		node.Next, right = right, node

		if levelLength--; levelLength == 0 {
			right, level, levelLength = nil, level+1, 1<<level
		}
	}

	return root
}
