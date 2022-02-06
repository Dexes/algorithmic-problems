package main

func countCharacters(words []string, chars string) int {
	result, counter, word := 0, count(chars, make([]byte, 123)), make([]byte, 123)
	for i := 0; i < len(words); i++ {
		if len(words[i]) > len(chars) {
			continue
		}

		if check(counter, count(words[i], word)) {
			result += len(words[i])
		}
	}

	return result
}

func check(a, b []byte) bool {
	for i := 'a'; i <= 'z'; i++ {
		if a[i] < b[i] {
			return false
		}
	}

	return true
}

func count(s string, counter []byte) []byte {
	for i := 'a'; i <= 'z'; i++ {
		counter[i] = 0
	}

	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	return counter
}
