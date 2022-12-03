package main

func findErrorNums(nums []int) []int {
	set := make([]bool, len(nums)+1)
	repetition, loss := 0, 0
	for i := 0; i < len(nums); i++ {
		if set[nums[i]] {
			repetition = nums[i]
		} else {
			set[nums[i]] = true
		}
	}

	for loss = 1; set[loss]; loss++ {
	}

	return []int{repetition, loss}
}
