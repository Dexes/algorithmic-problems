package main

func findDisappearedNumbers(nums []int) []int {
	set, i, number := makeBitset(len(nums)+1), 0, 0
	for ; i < len(nums); i++ {
		set.Set(nums[i])
	}

	for i, number = 0, 1; number <= len(nums); number++ {
		if !set.Test(number) {
			nums[i] = number
			i++
		}
	}

	return nums[:i]
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

func (b *Bitset) Test(n int) bool {
	return b.bits[n/64]&(1<<(n%64)) > 0
}
