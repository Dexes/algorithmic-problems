package main

func letterCasePermutation(s string) []string {
	return NewPermutator(s).Make()
}

type Permutator struct {
	permutations []string
	permutation  []byte
}

func NewPermutator(s string) *Permutator {
	return &Permutator{
		permutations: make([]string, 0, countLetters(s)<<1),
		permutation:  []byte(s),
	}
}

func (p *Permutator) Make() []string {
	p.make(0)
	return p.permutations
}

func (p *Permutator) make(index int) {
	if index == len(p.permutation) {
		p.permutations = append(p.permutations, string(p.permutation))
		return
	}

	p.make(index + 1)

	if p.permutation[index] > '9' {
		var b byte

		b, p.permutation[index] = p.permutation[index], invert(p.permutation[index])
		p.make(index + 1)
		p.permutation[index] = b
	}
}

func countLetters(s string) (result int) {
	for i := 0; i < len(s); i++ {
		if s[i] > '9' {
			result++
		}
	}

	return result
}

func invert(b byte) byte {
	if b < 'a' {
		return b + 32
	}

	return b - 32
}
