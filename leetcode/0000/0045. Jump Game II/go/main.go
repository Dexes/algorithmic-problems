package main

func jump(nums []int) int {
	dp, lastIndex := make([]int, len(nums)), len(nums)-1
	for i := 0; i < lastIndex; i++ {
		next := dp[i] + 1

		for j := min(lastIndex, i+nums[i]); j > i; j-- {
			if dp[j] == 0 || dp[j] > next {
				dp[j] = next
			}
		}
	}

	return dp[lastIndex]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
