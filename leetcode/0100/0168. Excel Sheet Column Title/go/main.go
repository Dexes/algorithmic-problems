package main

func convertToTitle(columnNumber int) string {
	result, index := make([]byte, 10), 9
	for columnNumber > 0 {
		columnNumber, result[index] = extractDigit(columnNumber)
		index--
	}

	return string(result[index+1:])
}

func extractDigit(number int) (int, byte) {
	reminder := byte(number % 26)
	if reminder == 0 {
		return (number - 1) / 26, 'Z'
	}

	return number / 26, 'A' - 1 + reminder
}
