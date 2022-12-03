package main

func isArmstrong(n int) bool {
	if n < 10 {
		return true
	}

	if n < 150 {
		return false
	}

	sum, length := 0, countDigits(n)
	for x := n; x > 0; x /= 10 {
		sum += pow(x%10, length)
	}

	return sum == n
}

func pow(n int, pow int) int {
	result := 1
	for pow > 0 {
		if pow&1 == 1 {
			result *= n
		}

		n *= n
		pow >>= 1
	}

	return result
}

func countDigits(n int) int {
	if n < 1e3 {
		return 3
	}

	if n < 1e4 {
		return 4
	}

	if n < 1e5 {
		return 5
	}

	if n < 1e6 {
		return 6
	}

	if n < 1e7 {
		return 7
	}

	if n < 1e8 {
		return 8
	}

	return 9
}
