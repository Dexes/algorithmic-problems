package main

func sumOfUnique(nums []int) int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	result := 0
	for num, count := range counter {
		if count == 1 {
			result += num
		}
	}

	return result
}
