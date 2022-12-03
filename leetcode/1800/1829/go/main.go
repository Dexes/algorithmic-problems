package main

func getMaximumXor(nums []int, maximumBit int) []int {
	result, xor, mask := make([]int, len(nums)), 0, 1<<maximumBit-1
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		xor ^= nums[i]
		result[j] = mask ^ xor
	}

	return result
}
