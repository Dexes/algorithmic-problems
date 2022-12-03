package main

func findPairs(nums []int, k int) int {
	if k == 0 {
		return countRepeating(toMap(nums))
	}

	return countPairs(toMap(nums), k)
}

func countPairs(counter map[int]int, k int) int {
	result := 0
	for num, _ := range counter {
		if counter[num+k] > 0 {
			result++
		}
	}

	return result
}

func countRepeating(counter map[int]int) int {
	result := 0
	for _, count := range counter {
		if count > 1 {
			result++
		}
	}

	return result
}

func toMap(nums []int) map[int]int {
	result := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		result[nums[i]]++
	}

	return result
}
