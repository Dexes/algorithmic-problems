package main

func check(nums []int) bool {
	i := 1
	for ; i < len(nums) && nums[i-1] <= nums[i]; i++ {
	}

	if i == len(nums) {
		return true
	}

	for i++; i < len(nums) && nums[i-1] <= nums[i]; i++ {
	}

	return i == len(nums) && nums[i-1] <= nums[0]
}
