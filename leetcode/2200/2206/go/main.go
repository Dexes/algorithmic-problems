package main

func divideArray(nums []int) bool {
	counter := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for _, count := range counter {
		if count&1 == 1 {
			return false
		}
	}

	return true
}
