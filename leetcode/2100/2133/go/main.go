package main

func checkValid(matrix [][]int) bool {
	bitset := makeBitset(len(matrix) * 2)

	for i := 0; i < len(matrix); i++ {
		bitset.Clear()

		for j := 0; j < len(matrix); j++ {
			bitset.Set(matrix[i][j] - 1)
			bitset.Set(matrix[j][i] - 1 + len(matrix))
		}

		if !bitset.IsFilledUp() {
			return false
		}
	}

	return true
}

type Bitset struct {
	bits   []uint64
	length int
}

func makeBitset(capacity int) *Bitset {
	return &Bitset{bits: make([]uint64, (capacity+63)/64), length: capacity}
}

func (b *Bitset) Set(n int) {
	b.bits[n/64] |= 1 << (n % 64)
}

func (b *Bitset) IsFilledUp() bool {
	length := b.length / 64
	for i := 0; i < length; i++ {
		if b.bits[i] != 0xffffffffffffffff {
			return false
		}
	}

	if b.length%64 > 0 {
		return b.bits[length] == 1<<(b.length%64)-1
	}

	return true
}

func (b *Bitset) Clear() {
	for i := 0; i < len(b.bits); i++ {
		b.bits[i] = 0
	}
}
