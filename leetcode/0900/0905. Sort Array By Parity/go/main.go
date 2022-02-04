package main

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for ; left < right && nums[left]&1 == 0; left++ {
		}

		for ; right > left && nums[right]&1 == 1; right-- {
		}

		nums[left], nums[right] = nums[right], nums[left]
		left, right = left+1, right-1
	}

	return nums
}
