package main

func sortArray(nums []int) []int {
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, left, right int) {
	if left >= right {
		return
	}

	x := left
	for i := left; i <= right; i++ {
		if nums[i] < nums[right] {
			nums[x], nums[i] = nums[i], nums[x]
			x++
		}
	}

	nums[x], nums[right] = nums[right], nums[x]

	sort(nums, left, x-1)
	sort(nums, x+1, right)
}
