package main

func kLengthApart(nums []int, k int) bool {
	i, j := 0, 0
	for ; i < len(nums) && nums[i] == 0; i++ {
	}

	for ; i < len(nums); i = j {
		for j = i + 1; j < len(nums) && nums[j] == 0; j++ {
		}

		if j < len(nums) && j-i <= k {
			return false
		}
	}

	return true
}
