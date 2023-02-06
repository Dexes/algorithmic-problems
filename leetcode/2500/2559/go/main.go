package main

func vowelStrings(words []string, queries [][]int) []int {
	prefix := toPrefix(words)
	result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		result[i] = prefix[queries[i][1]]
		if queries[i][0] > 0 {
			result[i] -= prefix[queries[i][0]-1]
		}
	}

	return result
}

func toPrefix(words []string) []int {
	prefix := make([]int, len(words))
	if check(words[0]) {
		prefix[0] = 1
	}

	for i := 1; i < len(words); i++ {
		prefix[i] = prefix[i-1]
		if check(words[i]) {
			prefix[i]++
		}
	}

	return prefix
}

func check(s string) bool {
	return isVowel(s[0]) && isVowel(s[len(s)-1])
}

func isVowel(char uint8) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}
