package main

func rearrangeArray(nums []int) []int {
	result := make([]int, len(nums))
	pIndex, nIndex := 0, 1

	for _, n := range nums {
		if n > 0 {
			result[pIndex], pIndex = n, pIndex+2
		} else {
			result[nIndex], nIndex = n, nIndex+2
		}
	}

	return result
}
