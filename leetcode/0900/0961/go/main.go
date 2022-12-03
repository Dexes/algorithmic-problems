package main

var set = makeBitset(10001)

func repeatedNTimes(nums []int) int {
	set.Clear()
	for i := 0; i < len(nums); i++ {
		if set.Test(nums[i]) {
			return nums[i]
		}

		set.Set(nums[i])
	}

	return -1
}

type Bitset []uint64

func makeBitset(capacity int) Bitset {
	return make(Bitset, (capacity+63)/64)
}

func (b Bitset) Set(n int) {
	b[n/64] |= 1 << (n % 64)
}

func (b Bitset) Test(n int) bool {
	return b[n/64]&(1<<(n%64)) > 0
}

func (b Bitset) Clear() {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}
