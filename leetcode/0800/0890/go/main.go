package main

var buffer = make([]int, 123)

func findAndReplacePattern(words []string, pattern string) []string {
	pIndices, wIndices, index := make([]int, len(pattern)), make([]int, len(pattern)), 0
	toIndices(pattern, pIndices)

	for i := 0; i < len(words); i++ {
		if toIndices(words[i], wIndices); equals(pIndices, wIndices) {
			words[index] = words[i]
			index++
		}
	}

	return words[:index]
}

func toIndices(s string, result []int) {
	for i := 'a'; i <= 'z'; i++ {
		buffer[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if buffer[s[i]] < 0 {
			buffer[s[i]] = i
		}

		result[i] = buffer[s[i]]
	}
}

func equals(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
