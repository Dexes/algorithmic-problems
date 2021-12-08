package main

func intersection(nums1 []int, nums2 []int) []int {
	return intersect(toSet(nums1), toSet(nums2))
}

func intersect(a map[int]bool, b map[int]bool) []int {
	result := make([]int, 0, min(len(a), len(b)))
	for k, _ := range a {
		if b[k] {
			result = append(result, k)
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

func toSet(a []int) map[int]bool {
	result := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		result[a[i]] = true
	}

	return result
}
