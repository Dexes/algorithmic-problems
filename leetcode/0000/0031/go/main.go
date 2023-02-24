package main

func nextPermutation(nums []int) {
	pivot := len(nums) - 2
	for ; pivot >= 0 && nums[pivot] >= nums[pivot+1]; pivot-- {
	}

	if pivot < 0 {
		reverse(nums)
		return
	}

	nextPivot := len(nums) - 1
	for ; nextPivot > pivot && nums[nextPivot] <= nums[pivot]; nextPivot-- {
	}

	nums[pivot], nums[nextPivot] = nums[nextPivot], nums[pivot]
	reverse(nums[pivot+1:])
}

func reverse(nums []int) {
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
