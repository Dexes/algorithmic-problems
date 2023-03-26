package main

import (
	"sort"
)

func minOperations(nums, queries []int) []int64 {
	sort.Ints(nums)
	prefix, last := toPrefix(nums), len(nums)-1

	result := make([]int64, len(queries))
	for i, q := range queries {
		if nums[0] >= q {
			result[i] = int64(prefix[last] - (q * len(nums)))
			continue
		}

		if nums[last] <= q {
			result[i] = int64((q * len(nums)) - prefix[last])
			continue
		}

		x := sort.Search(len(nums), func(i int) bool { return nums[i] > q })
		increase := (q * x) - prefix[x-1]
		decrease := prefix[last] - prefix[x-1] - (q * (len(nums) - x))
		result[i] = int64(increase + decrease)
	}

	return result
}

func toPrefix(nums []int) []int {
	result := make([]int, len(nums))
	result[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}

	return result
}
