package main

import "sort"

func intersection(nums [][]int) []int {
	result, index := nums[0], 0
	for num, cnt := range count(nums) {
		if cnt == len(nums) {
			result[index] = num
			index++
		}
	}

	sort.Ints(result[:index])
	return result[:index]
}

func count(nums [][]int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			result[nums[i][j]]++
		}
	}

	return result
}
