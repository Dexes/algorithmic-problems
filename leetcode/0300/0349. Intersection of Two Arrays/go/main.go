package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	set := make(map[int]struct{}, len(nums1))
	for i := 0; i < len(nums1); i++ {
		set[nums1[i]] = struct{}{}
	}

	wIndex := 0
	for i := 0; i < len(nums2); i++ {
		if _, ok := set[nums2[i]]; ok {
			nums1[wIndex] = nums2[i]
			wIndex++
			delete(set, nums2[i])
		}
	}

	return nums1[:wIndex]
}
