package main

func findDifference(nums1, nums2 []int) [][]int {
	a, b := convert(nums1), convert(nums2)
	return [][]int{remove(nums1, b), remove(nums2, a)}
}

func remove(nums []int, removed []bool) []int {
	wIndex := 0
	for i := 0; i < len(nums); i++ {
		if x := nums[i] + 1000; !removed[x] {
			removed[x] = true
			nums[wIndex] = nums[i]
			wIndex++
		}
	}

	return nums[:wIndex]
}

func convert(nums []int) []bool {
	result := make([]bool, 2001)
	for i := 0; i < len(nums); i++ {
		result[nums[i]+1000] = true
	}

	return result
}
