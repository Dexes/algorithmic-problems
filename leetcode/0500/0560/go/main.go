package main

func subarraySum(nums []int, k int) int {
	result, sum, sums := 0, 0, make(map[int]int)
	sums[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result += sums[sum-k]
		sums[sum]++
	}

	return result
}
