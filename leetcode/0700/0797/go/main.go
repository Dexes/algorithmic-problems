package main

func allPathsSourceTarget(graph [][]int) [][]int {
	return NewPermutator(graph).Make()
}

type Permutator struct {
	paths    [][]int
	path     []int
	lastNode int
	graph    [][]int
}

func NewPermutator(graph [][]int) *Permutator {
	return &Permutator{
		paths:    make([][]int, 0, 64),
		path:     make([]int, len(graph)),
		lastNode: len(graph) - 1,
		graph:    graph,
	}
}

func (p *Permutator) Make() [][]int {
	p.make(0, 0)
	return p.paths
}

func (p *Permutator) make(node, pathIndex int) {
	p.path[pathIndex] = node
	if node == p.lastNode {
		p.savePath(pathIndex + 1)
		return
	}

	for _, next := range p.graph[node] {
		p.make(next, pathIndex+1)
	}
}

func (p *Permutator) savePath(length int) {
	x := make([]int, length)
	copy(x, p.path)

	p.paths = append(p.paths, x)
}
