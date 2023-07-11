package main

func alternatingSubarray(nums []int) (result int) {
	for i := 0; i < len(nums); {
		current := count(nums, i)
		i += max(current-1, 1)

		if current > result {
			result = current
		}
	}

	if result < 2 {
		return -1
	}

	return result
}

func count(nums []int, index int) int {
	if index+1 >= len(nums) {
		return 1
	}

	current, next := nums[index], nums[index+1]
	if next-current != 1 {
		return 1
	}

	i := index + 2
	for ; i < len(nums) && nums[i] == current; i++ {
		current, next = next, current
	}

	return i - index
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
