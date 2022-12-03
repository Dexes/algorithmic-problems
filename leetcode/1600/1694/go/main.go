package main

func reformatNumber(number string) string {
	capacity := len(number)/3*4 + len(number)%3
	result, groupSize := make([]byte, 0, capacity), 0

	for i := findDigit(number, 0); i < len(number); i = findDigit(number, i+1) {
		result, groupSize = append(result, number[i]), groupSize+1
		if groupSize == 3 {
			result, groupSize = append(result, '-'), 0
		}
	}

	rightIndex := len(result) - 1
	if result[rightIndex] == '-' {
		return string(result[:rightIndex])
	}

	if result[rightIndex-1] == '-' {
		result[rightIndex-1], result[rightIndex-2] = result[rightIndex-2], result[rightIndex-1]
	}

	return string(result)
}

func findDigit(s string, index int) int {
	for ; index < len(s) && s[index] < '0'; index++ {
	}

	return index
}
