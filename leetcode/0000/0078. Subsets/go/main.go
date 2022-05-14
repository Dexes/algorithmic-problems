package main

func subsets(nums []int) [][]int {
	return NewPermutator(nums).Make()
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	numbers      []int
}

func NewPermutator(numbers []int) *Permutator {
	return &Permutator{
		permutations: make([][]int, 0, 1<<len(numbers)),
		permutation:  make([]int, len(numbers)),
		numbers:      numbers,
	}
}

func (p *Permutator) Make() [][]int {
	p.savePermutation(0)
	p.make(0, 0)
	return p.permutations
}

func (p *Permutator) make(nIndex, pIndex int) {
	if nIndex == len(p.numbers) {
		return
	}

	for ; nIndex < len(p.numbers); nIndex++ {
		p.permutation[pIndex] = p.numbers[nIndex]
		p.savePermutation(pIndex + 1)

		p.make(nIndex+1, pIndex+1)
	}
}

func (p *Permutator) savePermutation(length int) {
	x := make([]int, length)
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
