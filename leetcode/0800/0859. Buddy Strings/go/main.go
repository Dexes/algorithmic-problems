package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		return checkRepeated(s)
	}

	index, counter := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == goal[i] {
			continue
		}

		switch counter {
		case 0:
			index, counter = i, 1
		case 1:
			if s[index] != goal[i] || goal[index] != s[i] {
				return false
			}

			counter = 2
		default:
			return false
		}
	}

	return counter == 2
}

func checkRepeated(s string) bool {
	letters := make([]bool, 133)
	for i := 0; i < len(s); i++ {
		if letters[s[i]] {
			return true
		}

		letters[s[i]] = true
	}

	return false
}
