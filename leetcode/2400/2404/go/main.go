package main

func mostFrequentEven(nums []int) int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if nums[i]&1 == 1 {
			continue
		}

		counter[nums[i]]++
	}

	result, count := -1, 0
	for x, f := range counter {
		if f > count || f == count && x < result {
			result, count = x, f
		}
	}

	return result
}
