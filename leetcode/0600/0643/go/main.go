package main

func findMaxAverage(nums []int, k int) float64 {
	sum, count := 0, float64(k)
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	result := float64(sum) / count
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if float64(sum)/count > result {
			result = float64(sum) / count
		}
	}

	return result
}
