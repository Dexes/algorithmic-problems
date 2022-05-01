package main

func countPrefixes(words []string, s string) (result int) {
	for i := 0; i < len(words); i++ {
		if startsWith(s, words[i]) {
			result++
		}
	}

	return result
}

func startsWith(s, prefix string) bool {
	if len(prefix) > len(s) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if prefix[i] != s[i] {
			return false
		}
	}

	return true
}
