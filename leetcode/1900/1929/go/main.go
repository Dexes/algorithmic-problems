package main

func getConcatenation(nums []int) []int {
	length := len(nums)
	result := make([]int, 2*length)
	for i := 0; i < length; i++ {
		result[i], result[length+i] = nums[i], nums[i]
	}

	return result
}
