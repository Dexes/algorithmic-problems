package main

func minStartValue(nums []int) int {
	sum, minSum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < minSum {
			minSum = sum
		}
	}

	return 1 - minSum
}
