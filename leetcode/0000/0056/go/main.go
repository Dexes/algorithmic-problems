package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	current := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[current][1] >= intervals[i][0] {
			if intervals[i][1] > intervals[current][1] {
				intervals[current][1] = intervals[i][1]
			}

			continue
		}

		current++
		intervals[current] = intervals[i]
	}

	return intervals[:current+1]
}
