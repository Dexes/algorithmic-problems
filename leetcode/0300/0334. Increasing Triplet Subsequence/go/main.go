package main

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	if len(nums) == 3 {
		return nums[0] < nums[1] && nums[1] < nums[2]
	}

	for left, middle, i := getLeftAndMiddle(nums); i < len(nums); i++ {
		switch {
		case middle < nums[i]:
			return true
		case left < nums[i] && nums[i] < middle:
			middle = nums[i]
		case nums[i] < left:
			left = nums[i]
		}
	}

	return false
}

func getLeftAndMiddle(nums []int) (left int, middle int, index int) {
	for left, index = nums[0], 1; index < len(nums); index++ {
		if nums[index] > left {
			return left, nums[index], index + 1
		}

		if nums[index] < left {
			left = nums[index]
		}
	}

	return 0, 0, len(nums)
}
