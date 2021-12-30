package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	result, pow := float64(1), abs(n)
	for pow > 0 {
		if pow&1 == 1 {
			result *= x
		}

		x *= x
		pow >>= 1
	}

	if n < 0 {
		return 1 / result
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
