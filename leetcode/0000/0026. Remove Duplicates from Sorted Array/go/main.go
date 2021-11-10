package main

func removeDuplicates(nums []int) int {
	currentIndex := 0
	length := len(nums)
	if length < 2 {
		return length
	}

	for i := 1; i < length; i++ {
		if nums[currentIndex] != nums[i] {
			currentIndex++
			nums[currentIndex] = nums[i]
		}
	}

	return currentIndex + 1
}
