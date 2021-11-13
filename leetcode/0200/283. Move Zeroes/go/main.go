package main

func moveZeroes(nums []int) {
	current := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[current] = nums[i]
			current++
		}
	}

	for ; current < len(nums); current++ {
		nums[current] = 0
	}
}
