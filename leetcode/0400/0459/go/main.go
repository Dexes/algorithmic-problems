package main

func repeatedSubstringPattern(s string) bool {
	for i := len(s) / 2; i > 0; i-- {
		if check(s, i) {
			return true
		}
	}

	return false
}

func check(s string, patternLength int) bool {
	if len(s)%patternLength != 0 {
		return false
	}

	for i := patternLength; i < len(s); {
		for j := 0; j < patternLength; i, j = i+1, j+1 {
			if s[j] != s[i] {
				return false
			}
		}
	}

	return true
}
