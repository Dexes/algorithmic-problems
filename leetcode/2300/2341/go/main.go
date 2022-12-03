package main

func numberOfPairs(nums []int) []int {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	pairs, leftover := 0, 0
	for _, count := range counter {
		pairs += count / 2
		leftover += count % 2
	}

	return []int{pairs, leftover}
}
