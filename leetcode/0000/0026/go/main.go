package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}
