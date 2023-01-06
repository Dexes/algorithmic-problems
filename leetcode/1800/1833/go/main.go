package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	index := 0
	for ; index < len(costs) && costs[index] <= coins; index++ {
		coins -= costs[index]
	}

	return index
}
