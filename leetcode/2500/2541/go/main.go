package main

func minOperations(nums1, nums2 []int, k int) int64 {
	if k == 0 {
		return equal(nums1, nums2)
	}

	a, b := 0, 0
	for i := 0; i < len(nums1); i++ {
		switch diff := nums1[i] - nums2[i]; {
		case (diff % k) > 0:
			return -1
		case diff < 0:
			a -= diff
		case diff > 0:
			b += diff
		}
	}

	if a == b {
		return int64(a / k)
	}

	return -1
}

func equal(nums1, nums2 []int) int64 {
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return -1
		}
	}

	return 0
}
