package main

func subsetXORSum(nums []int) int {
	result := 0
	for mask := 1<<len(nums) - 1; mask > 0; mask-- {
		xor := 0
		for i := 0; i < len(nums); i++ {
			bit := 1 << i
			if mask&bit == bit {
				xor ^= nums[i]
			}
		}

		result += xor
	}

	return result
}
