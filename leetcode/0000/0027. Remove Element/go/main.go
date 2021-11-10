package main

func removeElement(nums []int, val int) int {
	currentIndex := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[currentIndex] = nums[i]
			currentIndex++
		}
	}

	return currentIndex
}
