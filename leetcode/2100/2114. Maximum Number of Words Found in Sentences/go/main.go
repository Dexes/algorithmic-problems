package main

func mostWordsFound(sentences []string) int {
	result := 1
	for i := 0; i < len(sentences); i++ {
		spacesCount := 0
		for j := 0; j < len(sentences[i]); j++ {
			if sentences[i][j] == ' ' {
				spacesCount++
			}
		}

		if spacesCount+1 > result {
			result = spacesCount + 1
		}
	}

	return result
}
