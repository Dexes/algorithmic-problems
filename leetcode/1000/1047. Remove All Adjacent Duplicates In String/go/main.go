package main

func removeDuplicates(s string) string {
	stack, stackIndex := make([]byte, len(s)), -1

	for i := 0; i < len(s); i++ {
		if stackIndex >= 0 && stack[stackIndex] == s[i] {
			stackIndex--
			continue
		}

		stackIndex++
		stack[stackIndex] = s[i]
	}

	return string(stack[:stackIndex+1])
}
