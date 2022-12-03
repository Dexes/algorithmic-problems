package main

func countHillValley(nums []int) int {
	result := 0
	for i, rightIndex := trim(nums); i <= rightIndex; {
		left, current := nums[i-1], nums[i]
		for i++; i <= rightIndex && nums[i] == nums[i-1]; i++ {
		}

		if left < current && current > nums[i] || left > current && current < nums[i] {
			result++
		}
	}

	return result
}

func trim(nums []int) (int, int) {
	i, j := 1, len(nums)-2
	for ; i < len(nums) && nums[i] == nums[i-1]; i++ {
	}

	for ; j >= i && nums[j] == nums[j+1]; j-- {
	}

	return i, j
}
