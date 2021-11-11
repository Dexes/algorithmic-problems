package main

func myAtoi(s string) int {
	var start, sign, result int

	start = skipWhitespaces(s)
	sign, start = getSign(s, start)
	result = extractNumber(s, start) * sign

	if result < -1<<31 {
		return -1 << 31
	}

	if result > 1<<31-1 {
		return 1<<31 - 1
	}

	return result
}

func extractNumber(s string, start int) int {
	result := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return result
		}

		next := result*10 + int(s[i]-'0')
		if next < result {
			return 1 << 32
		}

		result = next
	}

	return result
}

func skipWhitespaces(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			return i
		}
	}

	return 0
}

func getSign(s string, start int) (int, int) {
	if len(s) == 0 {
		return 1, start
	}

	if s[start] == '-' {
		return -1, start + 1
	}

	if s[start] == '+' {
		return 1, start + 1
	}

	return 1, start
}
