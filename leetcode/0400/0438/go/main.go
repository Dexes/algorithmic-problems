package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	sLetters, pLetters := countLetters(s[:len(p)]), countLetters(p)
	result := make([]int, 0, 10000)
	for i := len(p); i < len(s); i++ {
		if compare(sLetters, pLetters) {
			result = append(result, i-len(p))
		}

		sLetters[s[i-len(p)]]--
		sLetters[s[i]]++
	}

	if compare(sLetters, pLetters) {
		result = append(result, len(s)-len(p))
	}

	return result
}

func compare(first, second []int) bool {
	for i := 'a'; i <= 'z'; i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func countLetters(s string) []int {
	result := make([]int, 123)
	for i := 0; i < len(s); i++ {
		result[s[i]]++
	}

	return result
}
