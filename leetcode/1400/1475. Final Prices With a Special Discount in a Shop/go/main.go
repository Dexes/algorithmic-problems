package main

func finalPrices(prices []int) []int {
	stack, stackIndex := make([]int, len(prices)), 0
	for i := len(prices) - 1; i >= 0; i-- {
		for ; stackIndex > 0 && stack[stackIndex] > prices[i]; stackIndex-- {
		}

		stackIndex++
		stack[stackIndex] = prices[i]

		prices[i] -= stack[stackIndex-1]
	}

	return prices
}
