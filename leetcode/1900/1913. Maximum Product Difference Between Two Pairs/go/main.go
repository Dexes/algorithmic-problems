package main

func maxProductDifference(nums []int) int {
	firstMin, secondMin, firstMax, secondMax := 10000, 10000, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= firstMin {
			firstMin, secondMin = nums[i], firstMin
		} else if nums[i] <= secondMin {
			secondMin = nums[i]
		}

		if nums[i] >= firstMax {
			firstMax, secondMax = nums[i], firstMax
		} else if nums[i] >= secondMax {
			secondMax = nums[i]
		}
	}

	return firstMax*secondMax - firstMin*secondMin
}
