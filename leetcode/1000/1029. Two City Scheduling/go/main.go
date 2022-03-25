package main

import (
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})

	result, n := 0, len(costs)/2
	for i := 0; i < n; i++ {
		result += costs[i][0] + costs[i+n][1]
	}

	return result
}
