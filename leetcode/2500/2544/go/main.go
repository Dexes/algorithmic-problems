package main

func alternateDigitSum(n int) int {
	result, sign := 0, 1
	for ; n > 0; n /= 10 {
		result += sign * (n % 10)
		sign = -sign
	}

	if sign == 1 {
		return -result
	}

	return result
}
