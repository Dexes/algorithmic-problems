package main

func permuteUnique(nums []int) [][]int {
	return NewPermutator(nums).Make()
}

func count(numbers []int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		result[numbers[i]]++
	}

	return result
}

type Permutator struct {
	permutations [][]int
	permutation  []int
	numbers      map[int]int
}

func NewPermutator(numbers []int) *Permutator {
	return &Permutator{
		permutations: make([][]int, 0, 100),
		permutation:  make([]int, len(numbers)),
		numbers:      count(numbers),
	}
}

func (p *Permutator) Make() [][]int {
	p.make(0)
	return p.permutations
}

func (p *Permutator) make(index int) {
	if index == len(p.permutation) {
		p.savePermutation()
		return
	}

	for n, c := range p.numbers {
		if c == 0 {
			continue
		}

		p.permutation[index], p.numbers[n] = n, c-1
		p.make(index + 1)
		p.numbers[n] = c
	}
}

func (p *Permutator) savePermutation() {
	x := make([]int, len(p.permutation))
	copy(x, p.permutation)

	p.permutations = append(p.permutations, x)
}
