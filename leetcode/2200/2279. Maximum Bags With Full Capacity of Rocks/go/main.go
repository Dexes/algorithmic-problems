package main

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) (result int) {
	wIndex := 0
	for i := 0; i < len(capacity); i++ {
		if capacity[wIndex] = capacity[i] - rocks[i]; capacity[wIndex] > 0 {
			wIndex++
		}
	}

	sort.Ints(capacity[:wIndex])

	for i := 0; i < wIndex && capacity[i] <= additionalRocks; i++ {
		additionalRocks -= capacity[i]
		result++
	}

	return result + len(capacity) - wIndex
}
