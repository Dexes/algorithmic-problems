package main

func makeEqual(words []string) bool {
	letters := make([]int, 123)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			letters[words[i][j]]++
		}
	}

	for i := 'a'; i <= 'z'; i++ {
		if letters[i]%len(words) > 0 {
			return false
		}
	}

	return true
}
