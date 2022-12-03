package main

func combine(n, k int) [][]int {
	return NewPermutator(n, k).Make()
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	used         []bool
}

func NewPermutator(n, k int) *Permutator {
	return &Permutator{
		permutations: make([][]int, 0, 100),
		permutation:  make([]int, k),
		used:         make([]bool, n+1),
	}
}

func (p *Permutator) Make() [][]int {
	p.make(1, 0)
	return p.permutations
}

func (p *Permutator) make(n, index int) {
	if index == len(p.permutation) {
		p.savePermutation()
		return
	}

	for ; n < len(p.used); n++ {
		if !p.used[n] {
			p.permutation[index], p.used[n] = n, true
			p.make(n+1, index+1)
			p.used[n] = false
		}
	}
}

func (p *Permutator) savePermutation() {
	x := make([]int, len(p.permutation))
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
