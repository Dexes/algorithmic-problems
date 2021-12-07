package main

func targetIndices(nums []int, target int) []int {
	index, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < target {
			index++
		} else if nums[i] == target {
			count++
		}
	}

	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = index
		index++
	}

	return result
}
