package main

func minLength(s string) int {
	stack, stackIndex := make([]byte, len(s)), 0
	stack[stackIndex] = s[0]

	for i := 1; i < len(s); i++ {
		if stackIndex >= 0 && s[i] == 'B' && stack[stackIndex] == 'A' {
			stackIndex--
			continue
		}

		if stackIndex >= 0 && s[i] == 'D' && stack[stackIndex] == 'C' {
			stackIndex--
			continue
		}

		stackIndex++
		stack[stackIndex] = s[i]
	}

	return stackIndex + 1
}
