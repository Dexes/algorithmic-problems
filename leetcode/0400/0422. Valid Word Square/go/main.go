package main

func validWordSquare(words []string) bool {
	if len(words) != len(words[0]) {
		return false
	}

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if len(words) <= j || len(words[j]) <= i || words[i][j] != words[j][i] {
				return false
			}
		}
	}

	return true
}
