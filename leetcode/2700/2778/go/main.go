package main

func sumOfSquares(nums []int) int {
	result := nums[0] * nums[0]
	for i := 1; i < len(nums); i++ {
		if (len(nums) % (i + 1)) == 0 {
			result += nums[i] * nums[i]
		}
	}

	return result
}
