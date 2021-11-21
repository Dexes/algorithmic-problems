package main

func subtractProductAndSum(n int) int {
	digit, product, sum := 0, 1, 0
	for n > 0 {
		digit, n = n%10, n/10
		product *= digit
		sum += digit
	}

	return product - sum
}
