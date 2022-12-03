package main

func restoreIpAddresses(s string) []string {
	return NewPermutator(s).Make()
}

type Permutator struct {
	source       string
	permutations []string
	permutation  [4]string
}

func NewPermutator(source string) *Permutator {
	return &Permutator{
		source:       source,
		permutations: make([]string, 0, 20),
		permutation:  [4]string{},
	}
}

func (p *Permutator) Make() []string {
	p.make(0, 0)
	return p.permutations
}

func (p *Permutator) make(sIndex, pIndex int) {
	if sIndex == len(p.source) || pIndex == len(p.permutation) {
		if sIndex == len(p.source) && pIndex == len(p.permutation) {
			p.savePermutation()
		}

		return
	}

	source := p.source[sIndex:]
	if source[0] == '0' {
		p.permutation[pIndex] = "0"
		p.make(sIndex+1, pIndex+1)
		return
	}

	p.permutation[pIndex] = source[:1]
	p.make(sIndex+1, pIndex+1)

	if len(source) > 1 {
		p.permutation[pIndex] = source[:2]
		p.make(sIndex+2, pIndex+1)
	}

	if len(source) > 2 && toInt(source) <= 255 {
		p.permutation[pIndex] = source[:3]
		p.make(sIndex+3, pIndex+1)
	}
}

func (p *Permutator) savePermutation() {
	permutation := make([]byte, 0, 15)
	permutation = append(permutation, p.permutation[0]...)
	permutation = append(permutation, '.')
	permutation = append(permutation, p.permutation[1]...)
	permutation = append(permutation, '.')
	permutation = append(permutation, p.permutation[2]...)
	permutation = append(permutation, '.')
	permutation = append(permutation, p.permutation[3]...)

	p.permutations = append(p.permutations, string(permutation))
}

func toInt(s string) int {
	return int(s[0]-'0')*100 + int(s[1]-'0')*10 + int(s[2]-'0')
}
