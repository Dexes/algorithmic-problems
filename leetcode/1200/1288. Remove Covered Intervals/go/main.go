package main

import "sort"

func removeCoveredIntervals(intervals [][]int) (result int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	for i, j := 0, 0; i < len(intervals); i = j {
		for j = i + 1; j < len(intervals) && intervals[i][1] >= intervals[j][1]; j++ {
		}

		result++
	}

	return result
}
