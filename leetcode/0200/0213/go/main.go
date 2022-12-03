package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return max(robLine(nums[:len(nums)-1]), robLine(nums[1:]))
}

func robLine(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	counter := make([]int, len(nums))
	counter[0], counter[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(counter); i++ {
		counter[i] = max(nums[i]+counter[i-2], counter[i-1])
	}

	return counter[len(counter)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
