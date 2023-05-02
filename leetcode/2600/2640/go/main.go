package main

func findPrefixScore(nums []int) []int64 {
	result := append(make([]int64, 0, len(nums)), int64(nums[0]*2))
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}

		result = append(result, result[i-1]+int64(nums[i]+max))
	}

	return result
}
