package main

func zeroFilledSubarray(nums []int) int64 {
	result := 0
	for i, j := 0, 0; i < len(nums); i = j {
		for ; i < len(nums) && nums[i] != 0; i++ {
		}

		for j = i; j < len(nums) && nums[j] == 0; j++ {
		}

		result += (j - i + 1) * (j - i) / 2
	}

	return int64(result)
}
