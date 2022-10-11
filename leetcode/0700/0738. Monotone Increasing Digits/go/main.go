package main

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}

	digits, replaceIndex := toDigits(n), -1
	for i := 1; i < len(digits); i++ {
		if digits[i-1] < digits[i] {
			digits[i], replaceIndex = digits[i]-1, i-1
		}
	}

	for ; replaceIndex >= 0; replaceIndex-- {
		digits[replaceIndex] = 9
	}

	return toNumber(digits)
}

func toDigits(x int) []int {
	result, index := make([]int, 10), 0
	for ; x > 0; x /= 10 {
		result[index] = x % 10
		index++
	}

	return result[:index]
}

func toNumber(digits []int) (result int) {
	for i := len(digits) - 1; i >= 0; i-- {
		result = result*10 + digits[i]
	}

	return result
}
