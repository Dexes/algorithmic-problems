package main

type CombinationIterator struct {
	characters  string
	permutation []byte
	indices     []int
	index       int
	difference  int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	indices := make([]int, combinationLength)
	indices[0] = -1

	return CombinationIterator{
		characters:  characters,
		permutation: make([]byte, combinationLength),
		index:       0,
		indices:     indices,
		difference:  len(characters) - combinationLength,
	}
}

func (x *CombinationIterator) Next() string {
	indices, index := x.indices, x.index

	indices[index]++
	for index++; index < len(indices); index++ {
		indices[index] = indices[index-1] + 1
	}

	// tailLength = len(x.indices)-1 - index; indices[index] + tailLength == len(x.characters)-1
	// indices[index] + (len(x.indices)-1 - index) == len(x.characters)-1
	// indices[index] - index == len(x.characters) - len(x.indices)
	for index--; index >= 0 && indices[index]-index == x.difference; index-- {
	}

	x.index = index
	return x.String()
}

func (x *CombinationIterator) HasNext() bool {
	return x.index >= 0
}

func (x *CombinationIterator) String() string {
	result, characters, indices := x.permutation, x.characters, x.indices
	for i := 0; i < len(result); i++ {
		result[i] = characters[indices[i]]
	}

	return string(result)
}
