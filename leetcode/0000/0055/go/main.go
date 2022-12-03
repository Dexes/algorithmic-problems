package main

func canJump(nums []int) bool {
	reachedIndex, lastIndex := 0, len(nums)-1
	for i := 0; i < lastIndex; i++ {
		if i > reachedIndex {
			return false
		}

		if reachedIndex = max(reachedIndex, i+nums[i]); reachedIndex >= len(nums)-1 {
			return true
		}
	}

	return len(nums) == 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
