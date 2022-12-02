package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	letters1, letters2 := make([]int, 123), make([]int, 123)
	for i := 0; i < len(word1); i++ {
		letters1[word1[i]]++
		letters2[word2[i]]++
	}

	counter := make([]int, len(word1)+1)
	for i := 'a'; i <= 'z'; i++ {
		if (letters1[i] == 0) != (letters2[i] == 0) {
			return false
		}

		counter[letters1[i]]++
		counter[letters2[i]]--
	}

	for i := 'a'; i <= 'z'; i++ {
		if counter[letters1[i]] != 0 {
			return false
		}
	}

	return true
}
