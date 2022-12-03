package main

func countGoodSubstrings(s string) int {
	subLength := 3
	if len(s) < subLength {
		return 0
	}

	result, letters := 0, make([]byte, 123)
	for i := 0; i < subLength; i++ {
		letters[s[i]]++
	}

	if check(letters) {
		result++
	}

	for i := subLength; i < len(s); i++ {
		letters[s[i-subLength]]--
		letters[s[i]]++

		if check(letters) {
			result++
		}
	}

	return result
}

func check(letters []byte) bool {
	for i := 'a'; i <= 'z'; i++ {
		if letters[i] > 1 {
			return false
		}
	}

	return true
}
