package main

import "sort"

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) <= 1 {
		return 0
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	x, y, result := stockPrices[0][0]-stockPrices[1][0], stockPrices[0][1]-stockPrices[1][1], 1
	for i := 2; i < len(stockPrices); i++ {
		firstArea := x * (stockPrices[i-1][1] - stockPrices[i][1])
		secondArea := y * (stockPrices[i-1][0] - stockPrices[i][0])

		if firstArea != secondArea {
			x, y = stockPrices[i-1][0]-stockPrices[i][0], stockPrices[i-1][1]-stockPrices[i][1]
			result++
		}
	}

	return result
}
