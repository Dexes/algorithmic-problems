package main

func mostFrequent(nums []int, key int) int {
	numbers, lastIndex := make(map[int]int), len(nums)-1
	for i := 0; i < lastIndex; i++ {
		if nums[i] == key {
			numbers[nums[i+1]]++
		}
	}

	result, max := 0, 0
	for num, count := range numbers {
		if count > max {
			result, max = num, count
		}
	}

	return result
}
