package main

const maxResult = 2147483647

func divide(dividend int, divisor int) int {
	result, sign := 0, 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	dividend, divisor = abs(dividend), abs(divisor)
	for dividend >= divisor {
		counter := -1
		for x := divisor; x <= dividend; x += x {
			counter++
		}

		result += 1 << counter
		dividend -= divisor << counter
	}

	return min(maxResult, result*sign)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
