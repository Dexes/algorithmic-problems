package main

func solveNQueens(n int) [][]string {
	return NewPermutator(n).Make()
}

type Permutator struct {
	permutations [][]string
	permutation  [][]byte
	columns      []bool
	straights    []bool
	reverses     []bool
	size         int
}

func NewPermutator(size int) *Permutator {
	permutation := make([][]byte, size)
	for i := 0; i < size; i++ {
		permutation[i] = make([]byte, size)
		for j := 0; j < size; j++ {
			permutation[i][j] = '.'
		}
	}

	return &Permutator{
		permutations: make([][]string, 0, 100),
		permutation:  permutation,
		columns:      make([]bool, size),
		straights:    make([]bool, size<<1-1),
		reverses:     make([]bool, size<<1-1),
		size:         size,
	}
}

func (p *Permutator) Make() [][]string {
	p.make(0)
	return p.permutations
}

func (p *Permutator) make(rowIndex int) {
	if rowIndex == p.size {
		p.savePermutation()
		return
	}

	columns, straights, reverses := p.columns, p.straights, p.reverses
	columnIndex, straightIndex, reverseIndex := 0, p.size-rowIndex-1, rowIndex
	for ; columnIndex < p.size; columnIndex, straightIndex, reverseIndex = columnIndex+1, straightIndex+1, reverseIndex+1 {
		if columns[columnIndex] || straights[straightIndex] || reverses[reverseIndex] {
			continue
		}

		p.columns[columnIndex], p.straights[straightIndex], p.reverses[reverseIndex] = true, true, true
		p.permutation[rowIndex][columnIndex] = 'Q'
		p.make(rowIndex + 1)
		p.columns[columnIndex], p.straights[straightIndex], p.reverses[reverseIndex] = false, false, false
		p.permutation[rowIndex][columnIndex] = '.'
	}
}

func (p *Permutator) savePermutation() {
	strings, bytes := make([]string, p.size), p.permutation
	for i := 0; i < len(strings); i++ {
		strings[i] = string(bytes[i])
	}

	p.permutations = append(p.permutations, strings)
}
