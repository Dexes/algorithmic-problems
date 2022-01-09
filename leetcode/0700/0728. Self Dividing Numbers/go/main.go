package main

func selfDividingNumbers(left int, right int) []int {
	result, digits := make([]int, 0, 10000), getDigits(left, digitsNumber(right)+1)
	for ; left <= right; left, digits = increment(left, digits) {
		if isSelfDividingNumber(left, digits) {
			result = append(result, left)
		}
	}

	return result
}

func increment(number int, digits []int) (int, []int) {
	for i := len(digits) - 1; ; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			continue
		}

		digits[i]++
		return number + 1, digits
	}
}

func isSelfDividingNumber(number int, digits []int) bool {
	i := 0
	for ; digits[i] == 0; i++ {
	}

	for ; i < len(digits); i++ {
		if digits[i] == 0 || number%digits[i] > 0 {
			return false
		}
	}

	return true
}

func getDigits(num int, resultLen int) []int {
	result := make([]int, resultLen)
	for i := resultLen - 1; i >= 0 && num > 0; i-- {
		result[i] = num % 10
		num /= 10
	}

	return result
}

func digitsNumber(num int) int {
	result := 0
	for num > 0 {
		result++
		num /= 10
	}

	return result
}
