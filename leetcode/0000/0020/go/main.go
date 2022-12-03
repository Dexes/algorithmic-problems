package main

func isValid(s string) bool {
	stack, stackIndex := make([]byte, len(s)/2), 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			if stackIndex == len(stack) {
				return false
			}

			stack[stackIndex], stackIndex = s[i], stackIndex+1
			continue
		}

		if stackIndex--; stackIndex < 0 || s[i]-1 != stack[stackIndex] && s[i]-2 != stack[stackIndex] {
			return false
		}
	}

	return stackIndex == 0
}
