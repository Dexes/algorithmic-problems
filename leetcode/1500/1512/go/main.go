package main

func numIdenticalPairs(nums []int) int {
	result := 0
	counter := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		result += counter[nums[i]]
		counter[nums[i]]++
	}

	return result
}
