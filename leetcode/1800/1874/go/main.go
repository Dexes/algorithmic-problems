package main

import "sort"

func minProductSum(nums1 []int, nums2 []int) (result int) {
	sort.Ints(nums1)
	sort.Ints(nums2)

	for i, lastIndex := 0, len(nums1)-1; i < len(nums1); i++ {
		result += nums1[i] * nums2[lastIndex-i]
	}

	return result
}
