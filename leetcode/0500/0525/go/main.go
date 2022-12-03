package main

func findMaxLength(nums []int) int {
	result, sum, sums := 0, 0, make(map[int]int)
	sums[0] = -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}

		if index, ok := sums[sum]; ok {
			if i-index > result {
				result = i - index
			}
		} else {
			sums[sum] = i
		}
	}

	return result
}
