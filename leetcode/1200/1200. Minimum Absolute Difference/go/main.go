package main

import (
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	result, index, minDifference := make([][]int, len(arr)-1), 1, arr[1]-arr[0]
	result[0] = []int{arr[0], arr[1]}

	for i := 2; i < len(arr); i++ {
		difference := arr[i] - arr[i-1]

		if difference < minDifference {
			result[0][0], result[0][1] = arr[i-1], arr[i]
			minDifference, index = difference, 1
			continue
		}

		if difference == minDifference {
			if len(result[index]) == 0 {
				result[index] = []int{arr[i-1], arr[i]}
			} else {
				result[index][0], result[index][1] = arr[i-1], arr[i]
			}

			index++
		}
	}

	return result[:index]
}
