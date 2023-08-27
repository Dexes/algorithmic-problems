package main

import (
	"sort"
)

func countPairs(nums []int, target int) (result int) {
	sort.Ints(nums)

	for l, r := 0, len(nums); ; l++ {
		r = sort.Search(r, func(i int) bool { return nums[l]+nums[l+i] >= target }) - 1
		if r <= 0 {
			return result
		}

		result += r
	}
}
