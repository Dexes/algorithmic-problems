package main

func makeGood(s string) string {
	stack, stackIndex := make([]byte, len(s)), -1
	for i := 0; i < len(s); i++ {
		if stackIndex >= 0 && check(stack[stackIndex], s[i]) {
			stackIndex--
			continue
		}

		stackIndex++
		stack[stackIndex] = s[i]
	}

	return string(stack[:stackIndex+1])
}

func check(a, b byte) bool {
	return a+32 == b || a-32 == b
}
