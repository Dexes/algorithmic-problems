package main

func buyChoco(prices []int, money int) int {
	first, second := prices[0], prices[1]
	if first > second {
		first, second = second, first
	}

	for i := 2; i < len(prices); i++ {
		if prices[i] < first {
			first, second = prices[i], first
			continue
		}

		if prices[i] < second {
			second = prices[i]
		}
	}

	if x := first + second; x <= money {
		return money - x
	}

	return money
}
