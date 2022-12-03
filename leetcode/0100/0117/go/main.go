package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type QueueItem struct {
	Node  *Node
	Level int
}

type Queue []*QueueItem

func (q Queue) Push(items ...*QueueItem) Queue {
	for _, item := range items {
		if item.Node != nil {
			q = append(q, item)
		}
	}

	return q
}

func (q Queue) Pop() (*QueueItem, Queue) {
	return q[0], q[1:]
}

func (q Queue) IsEmpty() bool {
	return len(q) == 0
}

func makeQueue(root *Node, size int) Queue {
	result := make(Queue, 0, size)
	return result.Push(&QueueItem{root, 1})
}

func connect(root *Node) *Node {
	var (
		item, right *QueueItem = nil, &QueueItem{nil, 0}
		queue                  = makeQueue(root, 6000)
	)

	for !queue.IsEmpty() {
		item, queue = queue.Pop()
		queue = queue.Push(&QueueItem{item.Node.Right, item.Level + 1}, &QueueItem{item.Node.Left, item.Level + 1})
		if right.Level == item.Level {
			item.Node.Next = right.Node
		}

		right = item
	}

	return root
}
