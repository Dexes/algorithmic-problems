package main

func mergeAlternately(word1 string, word2 string) string {
	longestWord, shortestLen := word2, len(word1)
	if len(word2) > shortestLen {
		longestWord, shortestLen = word1, len(word2)
	}

	result := make([]byte, 0, len(word1)+len(word2))
	for i := 0; i < shortestLen; i++ {
		result = append(result, word1[i], word2[i])
	}

	return string(append(result, longestWord[shortestLen:]...))
}
