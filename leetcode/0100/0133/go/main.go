package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	result := map[*Node]*Node{node: {Val: node.Val}}
	queue := makeQueue(50)
	queue.Push(node)

	for !queue.IsEmpty() {
		src := queue.Pop()
		dst := result[src]

		dst.Neighbors = make([]*Node, len(src.Neighbors))
		for i := 0; i < len(src.Neighbors); i++ {
			if x, visited := result[src.Neighbors[i]]; visited {
				dst.Neighbors[i] = x
				continue
			}

			dst.Neighbors[i] = &Node{Val: src.Neighbors[i].Val}
			result[src.Neighbors[i]] = dst.Neighbors[i]
			queue.Push(src.Neighbors[i])
		}
	}

	return result[node]
}

type Queue struct {
	first     []*Node
	pushIndex int
	popIndex  int
}

func makeQueue(capacity int) *Queue {
	return &Queue{first: make([]*Node, capacity), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(first *Node) {
	q.first[q.pushIndex] = first
	q.pushIndex++

	if q.pushIndex >= len(q.first) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() *Node {
	first := q.first[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.first) {
		q.popIndex = 0
	}

	return first
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
