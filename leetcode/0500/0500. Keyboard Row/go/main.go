package main

var rows = []int{2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3, 3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3}

func findWords(words []string) []string {
	current := 0
	for i := 0; i < len(words); i++ {
		if check(words[i]) {
			words[current] = words[i]
			current++
		}
	}

	return words[:current]
}

func check(word string) bool {
	row := getLetterRow(word[0])
	for i := 1; i < len(word); i++ {
		if getLetterRow(word[i]) != row {
			return false
		}
	}

	return true
}

func getLetterRow(b byte) int {
	if 'A' <= b && b <= 'Z' {
		return rows[b-'A']
	}

	return rows[b-'a']
}
