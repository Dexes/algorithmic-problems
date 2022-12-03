package main

func maxProfit(prices []int) int {
	result, min := 0, 10000
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if prices[i]-min > result {
			result = prices[i] - min
		}
	}

	return result
}
