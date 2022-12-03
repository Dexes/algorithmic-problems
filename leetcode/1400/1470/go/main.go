package main

func shuffle(nums []int, n int) []int {
	result := make([]int, n*2)
	for i := 0; i < n; i++ {
		result[i*2], result[i*2+1] = nums[i], nums[n+i]
	}

	return result
}
