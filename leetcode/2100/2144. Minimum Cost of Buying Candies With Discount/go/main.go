package main

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)

	result := 0
	for i := len(cost) - 1; i > 0; i -= 3 {
		result += cost[i] + cost[i-1]
	}

	if len(cost)%3 == 1 {
		return result + cost[0]
	}

	return result
}
