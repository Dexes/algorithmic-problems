package main

import "sort"

const modulo int = 1e9 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	return max(horizontalCuts, h) * max(verticalCuts, w) % modulo
}

func max(list []int, limit int) int {
	sort.Ints(list)

	result := list[0]
	if x := limit - list[len(list)-1]; x > result {
		result = x
	}

	for i := 1; i < len(list); i++ {
		if x := list[i] - list[i-1]; x > result {
			result = x
		}
	}

	return result
}
