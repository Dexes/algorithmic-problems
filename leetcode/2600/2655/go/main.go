package main

import (
	"sort"
)

func findMaximalUncoveredRanges(n int, ranges [][]int) [][]int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })

	result, end := make([][]int, 0, 128), 0
	for i := 0; i < len(ranges); i++ {
		if ranges[i][0] > end {
			result = append(result, []int{end, ranges[i][0] - 1})
			end = ranges[i][1] + 1
			continue
		}

		if ranges[i][1] >= end {
			end = ranges[i][1] + 1
		}
	}

	if end < n {
		result = append(result, []int{end, n - 1})
	}

	return result
}
