package main

func calculateTime(keyboard string, word string) int {
	indices := make([]int, 123)
	for i := 0; i < len(keyboard); i++ {
		indices[keyboard[i]] = i
	}

	result, finger := 0, 0
	for i := 0; i < len(word); i++ {
		result += abs(finger - indices[word[i]])
		finger = indices[word[i]]
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
