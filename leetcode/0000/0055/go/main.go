package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	reachedIndex, lastIndex := 0, len(nums)-1
	for i := 0; ; i++ {
		if i > reachedIndex {
			return false
		}

		if reachedIndex = max(reachedIndex, i+nums[i]); reachedIndex >= lastIndex {
			return true
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
