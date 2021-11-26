package main

func countKDifference(nums []int, k int) int {
	counter := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	result := 0
	for i := k; i < len(counter); i++ {
		result += counter[i] * counter[i-k]
	}

	return result
}
