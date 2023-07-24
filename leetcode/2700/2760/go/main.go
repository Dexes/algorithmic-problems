package main

func longestAlternatingSubarray(nums []int, threshold int) (result int) {
	for i := 0; i < len(nums); {
		current := length(nums, i, threshold)
		i += max(current, 1)

		if current > result {
			result = current
		}
	}

	return result
}

func length(nums []int, index, threshold int) int {
	i, expected := index, 0
	for ; i < len(nums) && nums[i] <= threshold && nums[i]&1 == expected; i++ {
		expected = ^expected & 1
	}

	return i - index
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
