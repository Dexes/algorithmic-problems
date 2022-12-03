package main

func countConsistentStrings(allowed string, words []string) int {
	result, allowedMask := 0, toMask(allowed)
	for i := 0; i < len(words); i++ {
		if isConsistent(allowedMask, toMask(words[i])) {
			result++
		}
	}

	return result
}

func isConsistent(first, second uint32) bool {
	for i := 0; i < 26; i++ {
		if first&(1<<i) == 0 && second&(1<<i) > 0 {
			return false
		}
	}

	return true
}

func toMask(s string) uint32 {
	result := uint32(0)
	for i := 0; i < len(s); i++ {
		result |= 1 << (s[i] - 'a')
	}

	return result
}
