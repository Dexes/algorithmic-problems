package main

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i < 0 {
		reverse(nums)
		return
	}

	j := len(nums) - 1
	for ; j > i && nums[j] <= nums[i]; j-- {
	}

	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
