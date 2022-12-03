package main

func combinationSum3(k int, n int) [][]int {
	return NewPermutator(k, n).Make()
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	targetSum    int
}

func NewPermutator(capacity, targetSum int) *Permutator {
	return &Permutator{permutations: make([][]int, 0, 100), permutation: make([]int, capacity), targetSum: targetSum}
}

func (p *Permutator) Make() [][]int {
	p.make(1, 0, 0)
	return p.permutations
}

func (p *Permutator) make(n, currentSum, i int) {
	if currentSum == p.targetSum && i == len(p.permutation) {
		p.savePermutation()
		return
	}

	if i == len(p.permutation) || currentSum > p.targetSum {
		return
	}

	for ; n < 10; n++ {
		p.permutation[i] = n
		p.make(n+1, currentSum+n, i+1)
	}
}

func (p *Permutator) savePermutation() {
	x := make([]int, len(p.permutation))
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
