package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greaterMap := toMap(nums2)
	for i := 0; i < len(nums1); i++ {
		nums1[i] = greaterMap[nums1[i]]
	}

	return nums1
}

func toMap(nums []int) map[int]int {
	stack, stackIndex := make([]int, len(nums)), -1
	result := make(map[int]int)

	for i := len(nums) - 1; i >= 0; i-- {
		for ; stackIndex >= 0 && stack[stackIndex] < nums[i]; stackIndex-- {
		}

		if stackIndex == -1 {
			result[nums[i]] = -1
		} else {
			result[nums[i]] = stack[stackIndex]
		}

		stackIndex++
		stack[stackIndex] = nums[i]
	}

	return result
}
