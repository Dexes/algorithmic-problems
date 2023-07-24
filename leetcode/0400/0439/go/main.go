package main

func parseTernary(expression string) string {
	if expression[0] == 'T' {
		if expression[3] == ':' {
			return expression[2:3]
		}

		return parseTernary(expression[2:])
	}

	balance, i := 1, 3
	for ; balance > 0; i++ {
		if expression[i] == '?' {
			balance++
		} else if expression[i] == ':' {
			balance--
		}
	}

	if j := i + 1; len(expression) == j || expression[j] == ':' {
		return expression[i:j]
	}

	return parseTernary(expression[i:])
}
