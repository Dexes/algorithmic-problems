package main

import (
	"sort"
)

func maxDivScore(nums, divisors []int) int {
	sort.Ints(nums)
	divisors = compress(divisors)

	result, max := divisors[0], count(nums, divisors[0])
	for i := 1; i < len(divisors); i++ {
		if x := count(nums, divisors[i]); x > max || x == max && divisors[i] < result {
			result, max = divisors[i], x
		}
	}

	return result
}

func count(nums []int, divisor int) (result int) {
	for i := len(nums) - 1; i >= 0 && nums[i] >= divisor; i-- {
		if nums[i]%divisor == 0 {
			result++
		}
	}

	return result
}

func compress(nums []int) []int {
	sort.Ints(nums)
	wIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			nums[wIndex] = nums[i]
			wIndex++
		}
	}

	return nums[:wIndex]
}
