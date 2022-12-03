package main

func findLHS(nums []int) int {
	result, counter := 0, toMap(nums)
	for num, count := range counter {
		if nextCount, ok := counter[num+1]; ok && count+nextCount > result {
			result = count + nextCount
		}
	}

	return result
}

func toMap(nums []int) map[int]int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	return counter
}
