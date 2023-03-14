package main

import "sort"

func maxScore(nums []int) int {
	pSum, nSum, nLength := preprocessing(nums)
	if pSum == 0 {
		return 0
	}

	if pSum > nSum {
		return len(nums)
	}

	sort.Ints(nums[:nLength])

	index := 0
	for pSum -= nSum; pSum <= 0; pSum, index = pSum-nums[index], index+1 {
	}

	return len(nums) - index
}

func preprocessing(nums []int) (int, int, int) {
	pSum, nSum, nLength := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}

		if nums[i] > 0 {
			pSum += nums[i]
			continue
		}

		nSum -= nums[i]
		nums[nLength], nLength = nums[i], nLength+1
	}

	return pSum, nSum, nLength
}
