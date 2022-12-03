package main

func totalNQueens(n int) int {
	return NewPermutator(n).Count()
}

type Permutator struct {
	counter   int
	columns   []bool
	straights []bool
	reverses  []bool
	size      int
}

func NewPermutator(size int) *Permutator {
	return &Permutator{
		counter:   0,
		columns:   make([]bool, size),
		straights: make([]bool, size<<1-1),
		reverses:  make([]bool, size<<1-1),
		size:      size,
	}
}

func (p *Permutator) Count() int {
	p.make(0)
	return p.counter
}

func (p *Permutator) make(rowIndex int) {
	if rowIndex == p.size {
		p.counter++
		return
	}

	columns, straights, reverses := p.columns, p.straights, p.reverses
	columnIndex, straightIndex, reverseIndex := 0, p.size-rowIndex-1, rowIndex
	for ; columnIndex < p.size; columnIndex, straightIndex, reverseIndex = columnIndex+1, straightIndex+1, reverseIndex+1 {
		if columns[columnIndex] || straights[straightIndex] || reverses[reverseIndex] {
			continue
		}

		p.columns[columnIndex], p.straights[straightIndex], p.reverses[reverseIndex] = true, true, true
		p.make(rowIndex + 1)
		p.columns[columnIndex], p.straights[straightIndex], p.reverses[reverseIndex] = false, false, false
	}
}
