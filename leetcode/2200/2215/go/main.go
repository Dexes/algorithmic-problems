package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	first, second := toMap(nums1), toMap(nums2)
	return [][]int{sub(first, second, nums1), sub(second, first, nums2)}
}

func sub(a, b map[int]bool, result []int) []int {
	index := 0
	for x := range a {
		if !b[x] {
			result[index] = x
			index++
		}
	}

	return result[:index]
}

func toMap(nums []int) map[int]bool {
	result := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		result[nums[i]] = true
	}

	return result
}
