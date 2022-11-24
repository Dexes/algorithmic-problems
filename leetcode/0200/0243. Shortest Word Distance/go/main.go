package main

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	a, b, result := -1, -1, len(wordsDict)
	for i := 0; i < len(wordsDict); i++ {
		if wordsDict[i] == word1 {
			a = i
		} else if wordsDict[i] == word2 {
			b = i
		}

		if a >= 0 && b >= 0 {
			result = min(result, abs(a-b))
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
