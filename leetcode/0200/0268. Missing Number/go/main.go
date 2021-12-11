package main

func missingNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result ^= i ^ nums[i]
	}

	return result ^ len(nums)
}
