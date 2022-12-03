package main

func isPossible(nums []int) bool {
	lists := make([][]int, 2002)
	shortCount := 0

	for i := 0; i < len(nums); i++ {
		prev, current := nums[i]+1000, nums[i]+1001
		lastIndex := len(lists[prev]) - 1

		if lastIndex == -1 {
			lists[current] = insert(lists[current], 1)
			shortCount++
			continue
		}

		if lists[prev][lastIndex] == 2 {
			shortCount--
		}

		lists[current] = insert(lists[current], lists[prev][lastIndex]+1)
		lists[prev] = lists[prev][:lastIndex]
	}

	return shortCount == 0
}

func insert(nums []int, x int) []int {
	if cap(nums) == 0 {
		nums = make([]int, 0, 10)
	}

	nums = append(nums, x)
	for i := len(nums) - 1; i > 0 && nums[i] > nums[i-1]; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}

	return nums
}
