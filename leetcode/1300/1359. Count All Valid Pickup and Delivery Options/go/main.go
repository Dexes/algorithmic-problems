package main

func countOrders(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * i * (2*i - 1) % 1000000007
	}

	return result
}
