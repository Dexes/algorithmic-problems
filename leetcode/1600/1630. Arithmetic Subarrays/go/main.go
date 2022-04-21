package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result, buffer := make([]bool, len(l)), make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		left, right := l[i], r[i]+1
		copy(buffer[left:right], nums[left:right])

		result[i] = check(buffer[left:right])
	}

	return result
}

func check(nums []int) bool {
	sort.Ints(nums)
	for i, diff := 2, nums[0]-nums[1]; i < len(nums); i++ {
		if nums[i-1]-nums[i] != diff {
			return false
		}
	}

	return true
}
