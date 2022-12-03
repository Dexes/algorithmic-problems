package main

import (
	"sort"
)

func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	result, x := 0, 0

	if nums[0] > 1 {
		result, k = sum(0, nums[0], k)
	}

	for i := 1; i < len(nums) && k > 0; i++ {
		x, k = sum(nums[i-1], nums[i], k)
		result += x
	}

	last := nums[len(nums)-1]
	x, k = sum(last, last+k+1, k)
	return int64(result + x)
}

func sum(start, stop, k int) (int, int) {
	n := stop - start - 1
	switch {
	case n < 0:
		return 0, k
	case n > k:
		n = k
	}

	return (start*2 + n + 1) * n / 2, k - n
}
