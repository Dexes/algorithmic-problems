package main

func isPrefixString(s string, words []string) bool {
	for len(s) > 0 && len(words) > 0 && equal(s, words[0]) {
		s, words = s[len(words[0]):], words[1:]
	}

	return len(s) == 0
}

func equal(a, b string) bool {
	if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(b); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
