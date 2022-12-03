package main

func checkIfPangram(sentence string) bool {
	letters := make(map[byte]bool)
	for i := 0; i < len(sentence); i++ {
		letters[sentence[i]] = true
	}

	return len(letters) == 26
}
