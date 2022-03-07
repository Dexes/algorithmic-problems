package main

import "sort"

func findTheDistanceValue(first []int, second []int, d int) int {
	sort.Ints(second)

	result := 0
	for i := 0; i < len(first); i++ {
		if !checkRange(second, first[i]-d, first[i]+d) {
			result++
		}
	}

	return result
}

func checkRange(arr []int, a, b int) bool {
	left, right := 0, len(arr)-1
	for left <= right {
		switch middle := left + (right-left)/2; {
		case arr[middle] < a:
			left = middle + 1
		case arr[middle] > b:
			right = middle - 1
		default:
			return true
		}
	}

	return false
}
