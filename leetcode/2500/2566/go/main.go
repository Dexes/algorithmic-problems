package main

func minMaxDifference(num int) int {
	digits := split(num)
	return max(digits) - min(digits)
}

func min(digits []int) (result int) {
	for i := 1; i < len(digits); i++ {
		result *= 10

		if digits[i] != digits[0] {
			result += digits[i]
		}
	}

	return result
}

func max(digits []int) (result int) {
	i := 0
	for ; i < len(digits); i++ {
		result = result*10 + 9

		if digits[i] < 9 {
			break
		}
	}

	for j := i + 1; j < len(digits); j++ {
		if digits[j] == digits[i] {
			result = result*10 + 9
		} else {
			result = result*10 + digits[j]
		}
	}

	return result
}

func split(num int) []int {
	result, index := make([]int, 9), 9
	for ; num > 0; num /= 10 {
		index--
		result[index] = num % 10
	}

	return result[index:]
}
