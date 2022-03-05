package main

func maxAscendingSum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); {
		sum := nums[i]
		for i++; i < len(nums) && nums[i-1] < nums[i]; i++ {
			sum += nums[i]
		}

		if sum > result {
			result = sum
		}
	}

	return result
}
