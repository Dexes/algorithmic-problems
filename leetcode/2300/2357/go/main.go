package main

func minimumOperations(nums []int) int {
	counter := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		counter[nums[i]] = struct{}{}
	}

	result := len(counter)
	if _, ok := counter[0]; ok {
		result--
	}

	return result
}
