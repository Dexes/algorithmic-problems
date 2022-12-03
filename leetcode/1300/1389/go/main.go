package main

func createTargetArray(nums []int, index []int) []int {
	result := make([]int, 0, len(nums))
	for i := 0; i < len(index); i++ {
		result = insert(result, index[i], nums[i])
	}

	return result
}

func insert(nums []int, index, value int) []int {
	nums = append(nums, value)
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}

	nums[index] = value
	return nums
}
