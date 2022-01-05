package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := make([][]int, 0, 20000)
	for i := 0; i < len(nums); i = skip(nums, i) {
		for j := i + 1; j < len(nums); j = skip(nums, j) {
			target := -nums[i] - nums[j]
			if exists(nums, j+1, len(nums), target) {
				result = append(result, []int{nums[i], nums[j], target})
			}
		}
	}

	return result
}

func skip(nums []int, start int) int {
	i := start + 1
	for ; i < len(nums) && nums[i] == nums[start]; i++ {
	}

	return i
}

func exists(nums []int, left, right, target int) bool {
	for left < right {
		middle := left + (right-left)/2
		if target < nums[middle] {
			right = middle
			continue
		}

		if target > nums[middle] {
			left = middle + 1
			continue
		}

		return true
	}

	return false
}
