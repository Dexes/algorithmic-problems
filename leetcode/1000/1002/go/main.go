package main

func commonChars(words []string) []string {
	return intersect(countLetters(words))
}

func intersect(letters [][]int) []string {
	result := make([]string, 0, 1000)
	for i := byte(0); i < 26; i++ {
		for j := minLetterCount(i, letters); j > 0; j-- {
			result = append(result, string(i+'a'))
		}
	}

	return result
}

func minLetterCount(letter byte, letters [][]int) int {
	result := 10000
	for i := 0; i < len(letters); i++ {
		result = min(result, letters[i][letter])
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func countLetters(words []string) [][]int {
	result := make([][]int, len(words))
	for i := 0; i < len(words); i++ {
		result[i] = make([]int, 26)
		for j := 0; j < len(words[i]); j++ {
			result[i][words[i][j]-'a']++
		}
	}

	return result
}
