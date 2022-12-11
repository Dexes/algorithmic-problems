package main

type Allocator struct {
	addresses [1001][]int
	lengths   []int
}

func Constructor(n int) Allocator {
	return Allocator{addresses: [1001][]int{}, lengths: make([]int, n)}
}

func (a *Allocator) Allocate(size int, mID int) int {
	for i, reserved := 0, 0; i < len(a.lengths); {
		if a.lengths[i] > 0 {
			i += a.lengths[i]
			reserved = 0
			continue
		}

		if reserved++; reserved == size {
			i -= size - 1
			a.lengths[i] = size
			a.addresses[mID] = append(a.addresses[mID], i)
			return i
		}

		i++
	}

	return -1
}

func (a *Allocator) Free(mID int) int {
	if a.addresses[mID] == nil {
		return 0
	}

	result := 0
	for i := 0; i < len(a.addresses[mID]); i++ {
		address := a.addresses[mID][i]
		result += a.lengths[address]
		a.lengths[address] = 0
	}

	a.addresses[mID] = a.addresses[mID][:0]
	return result
}
