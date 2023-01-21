package main

func getConcatenation(nums []int) []int {
	result := make([]int, 2*len(nums))
	for i := 0; i < len(nums); i++ {
		result[i], result[len(nums)+i] = nums[i], nums[i]
	}

	return result
}
