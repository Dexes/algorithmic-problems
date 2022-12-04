package main

func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	for i := findSpace(sentence, 0); ; i = findSpace(sentence, i+1) {
		if i == len(sentence) {
			return true
		}

		if sentence[i-1] != sentence[i+1] {
			return false
		}
	}
}

func findSpace(s string, i int) int {
	for ; i < len(s) && s[i] != ' '; i++ {
	}

	return i
}
