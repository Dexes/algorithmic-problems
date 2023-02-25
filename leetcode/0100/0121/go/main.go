package main

func maxProfit(prices []int) int {
	result, min := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}

		if x := prices[i] - min; x > result {
			result = x
		}
	}

	return result
}
