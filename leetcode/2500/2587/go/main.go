package main

import "sort"

func maxScore(nums []int) int {
	length := len(nums)
	pSum, nSum, nums := preprocessing(nums)
	if pSum == 0 {
		return 0
	}

	if pSum > nSum {
		return length
	}

	sort.Ints(nums)

	index := 0
	for pSum -= nSum; pSum <= 0; pSum, index = pSum-nums[index], index+1 {
	}

	return length - index
}

func preprocessing(nums []int) (int, int, []int) {
	positive, negative, length := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		if nums[i] > 0 {
			positive += nums[i]
			continue
		}

		negative -= nums[i]
		nums[length], length = nums[i], length+1
	}

	return positive, negative, nums[:length]
}
