package main

import "sort"

func maxCoins(piles []int) (result int) {
	sort.Ints(piles)

	for i := len(piles) / 3; i < len(piles); i += 2 {
		result += piles[i]
	}

	return result
}
