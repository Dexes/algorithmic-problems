package main

import "math"

func isThree(n int) bool {
	if n == 1 {
		return false
	}

	x := int(math.Sqrt(float64(n)))
	return x*x == n && isPrime(x)
}

func isPrime(x int) bool {
	for i := int(math.Sqrt(float64(x))); i > 1; i-- {
		if x%i == 0 {
			return false
		}
	}

	return true
}
