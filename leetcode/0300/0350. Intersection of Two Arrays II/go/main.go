package main

func intersect(nums1 []int, nums2 []int) []int {
	set := toSet(nums1)
	result := make([]int, 0, min(len(nums1), len(nums2)))
	for i := 0; i < len(nums2); i++ {
		if set[nums2[i]] > 0 {
			set[nums2[i]]--
			result = append(result, nums2[i])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func toSet(a []int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(a); i++ {
		result[a[i]]++
	}

	return result
}
