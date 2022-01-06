package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	queue := makeQueue()
	queue.PushChildren(root)

	for {
		nodes := queue.GetAll()
		for left, right := 0, len(nodes)-1; left < right; left, right = left+1, right-1 {
			if nodes[left] == nil && nodes[right] == nil {
				continue
			}

			if nodes[left] == nil || nodes[right] == nil || nodes[left].Val != nodes[right].Val {
				return false
			}
		}

		if !queue.PushChildren(nodes...) {
			return true
		}
	}
}

type Queue struct {
	data      []*TreeNode
	pushIndex int
	popIndex  int
}

func makeQueue() *Queue {
	return &Queue{data: make([]*TreeNode, 1000), pushIndex: 0, popIndex: 0}
}

func (q *Queue) PushChildren(nodes ...*TreeNode) bool {
	success := false

	for i := 0; i < len(nodes); i++ {
		if nodes[i] == nil {
			continue
		}

		success = success || nodes[i].Left != nil || nodes[i].Right != nil
		q.Push(nodes[i].Left)
		q.Push(nodes[i].Right)
	}

	return success
}

func (q *Queue) Push(node *TreeNode) {
	q.data[q.pushIndex] = node
	q.pushIndex++

	if q.pushIndex >= len(q.data) {
		q.pushIndex = 0
	}
}

func (q *Queue) GetAll() []*TreeNode {
	result := q.data[q.popIndex:q.pushIndex]
	q.popIndex = q.pushIndex

	return result
}
