package main

func countSubarrays(nums []int) (result int64) {
	for i, j := 0, 0; i < len(nums); i = j {
		for j++; j < len(nums) && nums[j] > nums[j-1]; j++ {
		}

		length := int64(j - i)
		result += length * (length + 1) / 2
	}

	return result
}
