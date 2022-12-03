package main

type Bitset struct {
	bits    []uint64
	length  int
	count   int
	flipped bool
}

func Constructor(size int) Bitset {
	return Bitset{bits: make([]uint64, (size+63)/64), length: size, count: 0, flipped: false}
}

func (b *Bitset) Fix(idx int) {
	if b.flipped {
		b.unfix(idx)
	} else {
		b.fix(idx)
	}
}

func (b *Bitset) Unfix(idx int) {
	if b.flipped {
		b.fix(idx)
	} else {
		b.unfix(idx)
	}
}

func (b *Bitset) Flip() {
	b.flipped = !b.flipped
}

func (b *Bitset) All() bool {
	if b.flipped {
		return b.count == 0
	}

	return b.count == b.length
}

func (b *Bitset) One() bool {
	if b.flipped {
		return b.count < b.length
	}

	return b.count > 0
}

func (b *Bitset) Count() int {
	if b.flipped {
		return b.length - b.count
	}

	return b.count
}

func (b *Bitset) ToString() string {
	result := make([]byte, 0, b.length)
	zero, one := byte('0'), byte('1')
	if b.flipped {
		zero, one = one, zero
	}

	length := b.length / 64
	for i := 0; i < length; i++ {
		result = b.fillString(result, b.bits[i], 64, zero, one)
	}

	if reminder := b.length % 64; reminder > 0 {
		result = b.fillString(result, b.bits[length], reminder, zero, one)
	}

	return string(result)
}

func (b *Bitset) fix(idx int) {
	index, bit := idx/64, uint64(1<<(idx%64))
	if b.bits[index]&bit == 0 {
		b.bits[index] |= bit
		b.count++
	}
}

func (b *Bitset) unfix(idx int) {
	index, bit := idx/64, uint64(1<<(idx%64))
	if b.bits[index]&bit > 0 {
		b.bits[index] ^= bit
		b.count--
	}
}

func (b *Bitset) fillString(bits []byte, number uint64, stop int, zero, one byte) []byte {
	for i := 0; i < stop; i++ {
		if number&(1<<i) == 0 {
			bits = append(bits, zero)
		} else {
			bits = append(bits, one)
		}
	}

	return bits
}
