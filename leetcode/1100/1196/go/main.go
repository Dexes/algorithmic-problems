package main

import "sort"

const max = 5000

func maxNumberOfApples(weight []int) int {
	sort.Ints(weight)

	for i := 1; i < len(weight); i++ {
		if weight[i] += weight[i-1]; weight[i] > max {
			return i
		}
	}

	return len(weight)
}
