package main

import "sort"

func maxPrice(items [][]int, capacity int) float64 {
	sort.Slice(items, func(i, j int) bool {
		return (float64(items[i][0]) / float64(items[i][1])) > (float64(items[j][0]) / float64(items[j][1]))
	})

	for price, i := 0, 0; ; i++ {
		if i == len(items) {
			return -1
		}

		if items[i][1] >= capacity {
			return float64(price) + (float64(items[i][0]) / (float64(items[i][1]) / float64(capacity)))
		}

		price += items[i][0]
		capacity -= items[i][1]
	}
}
