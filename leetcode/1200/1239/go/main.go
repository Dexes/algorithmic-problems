package main

func maxLength(arr []string) int {
	return NewPermutator(arr).GetMaxLength()
}

type Permutator struct {
	maxLength int
	masks     []StringMask
	usedMask  int
}

func NewPermutator(strings []string) *Permutator {
	return &Permutator{
		maxLength: 0,
		masks:     toMasks(strings),
		usedMask:  0,
	}
}

func (p *Permutator) GetMaxLength() int {
	p.calculateMaxLength(0, 0)
	return p.maxLength
}

func (p *Permutator) calculateMaxLength(index, length int) {
	if length > p.maxLength {
		p.maxLength = length
	}

	if index == len(p.masks) {
		return
	}

	var usedMask, currentMask int
	for ; index < len(p.masks); index++ {
		if usedMask, currentMask = p.usedMask, p.masks[index].Mask; (usedMask & currentMask) > 0 {
			continue
		}

		p.usedMask = usedMask | currentMask
		p.calculateMaxLength(index+1, length+p.masks[index].Length)
		p.usedMask = usedMask
	}
}

func toMasks(strings []string) []StringMask {
	result := make([]StringMask, 0, len(strings))
	for i := 0; i < len(strings); i++ {
		if mask, ok := toMask(strings[i]); ok {
			result = append(result, mask)
		}
	}

	return result
}

func toMask(s string) (StringMask, bool) {
	result := StringMask{Length: len(s)}
	for i := 0; i < len(s); i++ {
		bit := 1 << (s[i] - 'a')
		if (result.Mask & bit) > 0 {
			return result, false
		}

		result.Mask |= bit
	}

	return result, true
}

type StringMask struct {
	Length int
	Mask   int
}
