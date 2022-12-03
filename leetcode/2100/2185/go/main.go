package main

func prefixCount(words []string, prefix string) int {
	result := 0
	for i := 0; i < len(words); i++ {
		if check(words[i], prefix) {
			result++
		}
	}

	return result
}

func check(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}

	return true
}
