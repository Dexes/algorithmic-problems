package main

func findMatrix(nums []int) [][]int {
	result := make([][]int, 0, 8)
	counter, rIndex := make([]int, len(nums)+1), 0
	for _, n := range nums {
		rIndex = counter[n]
		counter[n]++

		if rIndex == len(result) {
			result = append(result, make([]int, 0, 4))
		}

		result[rIndex] = append(result[rIndex], n)
	}

	return result
}
