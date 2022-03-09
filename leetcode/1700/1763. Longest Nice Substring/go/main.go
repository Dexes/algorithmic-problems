package main

func longestNiceSubstring(s string) string {
	switch {
	case len(s) < 2:
		return ""
	case len(s) == 2:
		if s[0]+32 == s[1] || s[0] == s[1]+32 {
			return s
		}

		return ""
	}

	hindering, isEmpty := hinderingLetters(s)
	if isEmpty {
		return s
	}

	result, start := "", 0
	for i := 0; i < len(s); i++ {
		if hindering[s[i]] {
			result = longest(result, longestNiceSubstring(s[start:i]))
			start = i + 1
		}
	}

	return longest(result, longestNiceSubstring(s[start:]))
}

func longest(a, b string) string {
	if len(a) >= len(b) {
		return a
	}

	return b
}

func hinderingLetters(s string) ([]bool, bool) {
	letters := make([]bool, 133)
	for i := 0; i < len(s); i++ {
		letters[s[i]] = true
	}

	isEmpty := true
	for i := byte('A'); i <= 'Z'; i++ {
		switch {
		case letters[i] && letters[i+32]:
			letters[i], letters[i+32] = false, false
		case letters[i] || letters[i+32]:
			isEmpty = false
		}
	}

	return letters, isEmpty
}
