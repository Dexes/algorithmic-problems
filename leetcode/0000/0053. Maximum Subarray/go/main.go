package main

func maxSubArray(nums []int) int {
	sum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(sum+nums[i], nums[i])
		maxSum = max(maxSum, sum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
