package main

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	for i, j, stackIndex := 0, -1, -1; i < len(popped); i++ {
		if stackIndex >= 0 && popped[i] == pushed[stackIndex] {
			stackIndex--
			continue
		}

		for j++; j < len(pushed) && popped[i] != pushed[j]; j++ {
			stackIndex++
			pushed[stackIndex] = pushed[j]
		}

		if j == len(pushed) {
			return false
		}
	}

	return true
}
