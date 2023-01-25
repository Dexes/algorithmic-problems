package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	result := 1
	for i := 0; i < len(coins) && coins[i] <= result; i++ {
		result += coins[i]
	}

	return result
}
