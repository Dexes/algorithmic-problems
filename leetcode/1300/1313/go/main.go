package main

func decompressRLElist(nums []int) []int {
	result := make([]int, 0, 5000)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			result = append(result, nums[i+1])
		}
	}

	return result
}
