package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)

	fivePercent, sum := len(arr)/20, 0
	arr = arr[fivePercent : len(arr)-fivePercent]
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return float64(sum) / float64(len(arr))
}
