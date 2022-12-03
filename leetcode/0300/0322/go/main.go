package main

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	var exists bool
	if coins, exists = filterCoins(coins, amount); exists {
		return 1
	}

	coinsCounter := makeCoinsCounter(coins, amount)
	for i := 1; i < len(coinsCounter); i++ {
		if coinsCounter[i] == 0 {
			continue
		}

		for j, nextCount := 0, coinsCounter[i]+1; j < len(coins); j++ {
			nextAmount := i + coins[j]
			if nextAmount <= amount && (coinsCounter[nextAmount] == 0 || coinsCounter[nextAmount] > nextCount) {
				coinsCounter[nextAmount] = nextCount
			}
		}
	}

	if coinsCounter[amount] == 0 {
		return -1
	}

	return coinsCounter[amount]
}

func filterCoins(coins []int, amount int) ([]int, bool) {
	wIndex := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] == amount {
			return nil, true
		}

		if coins[i] < amount {
			coins[wIndex] = coins[i]
			wIndex++
		}
	}

	return coins[:wIndex], false
}

func makeCoinsCounter(coins []int, amount int) []int {
	result := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		if coins[i] <= amount {
			result[coins[i]] = 1
		}
	}

	return result
}
