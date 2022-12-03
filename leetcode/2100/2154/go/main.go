package main

func findFinalValue(nums []int, original int) int {
	set := toSet(nums)
	for original <= 1000 && set.Test(original) {
		original <<= 1
	}

	return original
}

func toSet(nums []int) *Bitset {
	result := makeBitset(1001)
	for i := 0; i < len(nums); i++ {
		result.Set(nums[i])
	}

	return result
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
