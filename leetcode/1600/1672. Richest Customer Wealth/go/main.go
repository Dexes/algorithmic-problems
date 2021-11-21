package main

func maximumWealth(accounts [][]int) int {
	result := 0
	for i := 0; i < len(accounts); i++ {
		result = max(result, sum(accounts[i]))
	}

	return result
}

func sum(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
