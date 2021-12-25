package main

func evalRPN(tokens []string) int {
	stack := make([]int, 0, 5000)
	for i := 0; i < len(tokens); i++ {
		if len(tokens[i]) == 1 && isOperator(tokens[i][0]) {
			value := calculate(tokens[i][0], stack[len(stack)-2], stack[len(stack)-1])
			stack = append(stack[:len(stack)-2], value)
			continue
		}

		stack = append(stack, toInt(tokens[i]))
	}

	return stack[0]
}

func calculate(operator byte, a, b int) int {
	switch operator {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		return a / b
	}
}

func isOperator(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/'
}

func toInt(s string) int {
	result, index, isNegative := 0, 0, false
	if s[0] == '-' {
		index, isNegative = 1, true
	}

	for ; index < len(s); index++ {
		result = result*10 + int(s[index]-'0')
	}

	if isNegative {
		return -result
	}

	return result
}
