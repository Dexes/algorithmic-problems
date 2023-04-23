package main

func sumOfMultiples(n int) int {
	return sum(3, n) + sum(5, n) + sum(7, n) + sum(105, n) - sum(15, n) - sum(21, n) - sum(35, n)
}

func sum(x, max int) int {
	max -= max % x
	return (max / x) * (x + max) / 2
}
