package main

func findKDistantIndices(nums []int, key int, k int) []int {
	result, maxResult, i, j := make([]int, 0, len(nums)), -1, -1, 0
	for {
		for i++; i < len(nums) && nums[i] != key; i++ {
		}

		if i == len(nums) {
			return result
		}

		j, maxResult = max(maxResult+1, i-k), min(len(nums)-1, i+k)
		for ; j <= maxResult; j++ {
			result = append(result, j)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
