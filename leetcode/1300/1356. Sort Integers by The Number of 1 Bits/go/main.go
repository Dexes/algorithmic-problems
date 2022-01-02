package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		iBits, jBits := countBits(arr[i]), countBits(arr[j])
		if iBits == jBits {
			return arr[i] < arr[j]
		}

		return iBits < jBits
	})

	return arr
}

func countBits(n int) int {
	result := 0
	for n > 0 {
		if n&1 == 1 {
			result++
		}

		n >>= 1
	}

	return result
}
