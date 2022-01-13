package main

import (
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sortPoints(points)

	result := 0
	for i := 0; i < len(points); {
		result++
		interval := points[i]

		for i++; i < len(points) && intersect(interval, points[i]); i++ {
			if points[i][0] > interval[0] {
				interval[0] = points[i][0]
			}

			if points[i][1] < interval[1] {
				interval[1] = points[i][1]
			}
		}
	}

	return result
}

func intersect(first, second []int) bool {
	return second[0] <= first[1] && first[1] <= second[1] ||
		first[0] <= second[0] && second[0] <= first[1]
}

func sortPoints(points [][]int) {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}

		return points[i][0] < points[j][0]
	})
}
