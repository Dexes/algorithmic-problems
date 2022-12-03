package main

func countWords(words1 []string, words2 []string) int {
	first, second := toMap(words1), toMap(words2)
	result := 0
	for k, v := range first {
		if v == 1 && second[k] == 1 {
			result++
		}
	}

	return result
}

func toMap(a []string) map[string]int {
	result := make(map[string]int)
	for i := 0; i < len(a); i++ {
		result[a[i]]++
	}

	return result
}
