package main

import (
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	distances := make([][]int, len(points))
	for i := 0; i < len(points); i++ {
		distances[i] = []int{points[i][0]*points[i][0] + points[i][1]*points[i][1], i}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i][0] < distances[j][0]
	})

	for i := 0; i < k; i++ {
		distances[i] = points[distances[i][1]]
	}

	return distances[:k]
}
