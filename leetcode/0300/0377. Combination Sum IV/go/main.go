package main

func combinationSum4(nums []int, target int) int {
	var dp []int
	dp, nums = makeDP(nums, target)

	for current := 2; current < len(dp); current++ {
		for i := 0; i < len(nums); i++ {
			if current >= nums[i] {
				dp[current] += dp[current-nums[i]]
			}
		}
	}

	return dp[target]
}

func makeDP(nums []int, target int) ([]int, []int) {
	result := make([]int, target+1)
	rIndex, wIndex := 0, 0
	for ; rIndex < len(nums); rIndex++ {
		if nums[rIndex] > target {
			continue
		}

		result[nums[rIndex]] = 1
		nums[wIndex] = nums[rIndex]
		wIndex++
	}

	return result, nums[:wIndex]
}
