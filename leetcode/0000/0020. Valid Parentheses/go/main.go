package main

type parenthese byte

func (p parenthese) IsOpen() bool {
	return p == '(' || p == '[' || p == '{'
}

func (p parenthese) IsClose(close parenthese) bool {
	return close == '(' && p == ')' ||
		close == '[' && p == ']' ||
		close == '{' && p == '}'
}

func isValid(s string) bool {
	stack := make([]parenthese, 0)

	for _, char := range s {
		p := parenthese(char)
		if p.IsOpen() {
			stack = append(stack, p)
			continue
		}

		lastIndex := len(stack) - 1
		if lastIndex == -1 || !p.IsClose(stack[lastIndex]) {
			return false
		}

		stack = stack[:lastIndex]
	}

	return len(stack) == 0
}
