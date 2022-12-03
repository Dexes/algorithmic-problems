package main

func anagramMappings(nums1 []int, nums2 []int) []int {
	indices := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		indices[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		nums1[i] = indices[nums1[i]]
	}

	return nums1
}
