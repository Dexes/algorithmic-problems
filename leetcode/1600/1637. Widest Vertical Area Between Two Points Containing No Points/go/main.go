package main

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := points[1][0] - points[0][0]
	for i := 2; i < len(points); i++ {
		if x := points[i][0] - points[i-1][0]; x > result {
			result = x
		}
	}

	return result
}
