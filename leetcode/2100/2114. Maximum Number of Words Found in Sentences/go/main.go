package main

func mostWordsFound(sentences []string) int {
	result := 0
	for i := 0; i < len(sentences); i++ {
		spacesCount := 0
		for j := 0; j < len(sentences[i]); j++ {
			if sentences[i][j] == ' ' {
				spacesCount++
			}
		}

		if spacesCount > result {
			result = spacesCount
		}
	}

	return result + 1
}
