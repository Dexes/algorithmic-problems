package main

func calculate(s string) int {
	result, _ := calculateInsideParentheses(s, 0)
	return result
}

func calculateInsideParentheses(s string, pos int) (int, int) {
	result, number, sign := 0, 0, 1

	for {
		pos = skipSpaces(s, pos)
		if pos == len(s) || s[pos] == ')' {
			return result, pos + 1
		}

		if s[pos] == '+' {
			sign = 1
			pos++
			continue
		}

		if s[pos] == '-' {
			sign = -1
			pos++
			continue
		}

		if s[pos] == '(' {
			number, pos = calculateInsideParentheses(s, pos+1)
		} else {
			number, pos = readInt(s, pos)
		}

		result += sign * number
	}
}

func readInt(s string, pos int) (int, int) {
	result := 0
	for ; pos < len(s) && isDigit(s[pos]); pos++ {
		result = result*10 + int(s[pos]-'0')
	}

	return result, pos
}

func skipSpaces(s string, pos int) int {
	for ; pos < len(s) && s[pos] == ' '; pos++ {
	}

	return pos
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
