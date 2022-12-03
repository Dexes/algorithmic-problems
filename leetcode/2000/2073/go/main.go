package main

func timeRequiredToBuy(tickets []int, k int) int {
	result := tickets[k]
	for i := 0; i < k; i++ {
		result += min(tickets[i], tickets[k])
	}

	tickets[k]--
	for i := k + 1; i < len(tickets); i++ {
		result += min(tickets[i], tickets[k])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
