package main

func validPath(n int, edges [][]int, start int, end int) bool {
	graph, visited, queue := toGraph(edges, n), makeBitset(n), makeQueue()
	queue.Push(start)

	for !queue.IsEmpty() {
		vertex := queue.Pop()
		if visited.Test(vertex) {
			continue
		}

		if vertex == end {
			return true
		}

		visited.Set(vertex)
		for i := 0; i < len(graph[vertex]); i++ {
			queue.Push(graph[vertex][i])
		}
	}

	return false
}

func toGraph(edges [][]int, n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < len(edges); i++ {
		result[edges[i][0]] = append(result[edges[i][0]], edges[i][1])
		result[edges[i][1]] = append(result[edges[i][1]], edges[i][0])
	}

	return result
}

type Bitset struct {
	bits   []uint64
	length int
}

func makeBitset(length int) *Bitset {
	return &Bitset{bits: make([]uint64, (length+63)/64), length: length}
}

func (b *Bitset) Set(n int) {
	b.bits[n/64] |= 1 << (n % 64)
}

func (b *Bitset) Test(n int) bool {
	return b.bits[n/64]&(1<<(n%64)) > 0
}

type Queue struct {
	data      []int
	pushIndex int
	popIndex  int
}

func makeQueue() *Queue {
	return &Queue{data: make([]int, 2000), pushIndex: 0, popIndex: 0}
}

func (q *Queue) Push(n int) {
	q.data[q.pushIndex] = n
	q.pushIndex++

	if q.pushIndex >= len(q.data) {
		q.pushIndex = 0
	}
}

func (q *Queue) Pop() int {
	n := q.data[q.popIndex]
	q.popIndex++

	if q.popIndex >= len(q.data) {
		q.popIndex = 0
	}

	return n
}

func (q *Queue) IsEmpty() bool {
	return q.pushIndex == q.popIndex
}
