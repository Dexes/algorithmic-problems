package main

func uncommonFromSentences(s1 string, s2 string) []string {
	first, second := toMap(s1), toMap(s2)
	result := make([]string, 0, len(first)+len(second))
	result = appendUnique(result, first, second)
	result = appendUnique(result, second, first)

	return result
}

func appendUnique(result []string, first, second map[string]int) []string {
	for word, count := range first {
		if count == 1 && second[word] == 0 {
			result = append(result, word)
		}
	}

	return result
}

func toMap(s string) map[string]int {
	result := make(map[string]int)
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result[s[start:i]]++
			start = i + 1
		}
	}

	result[s[start:]]++
	return result
}
