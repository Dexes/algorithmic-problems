package main

func scoreOfParentheses(s string) (result int) {
	for i := 0; i < len(s); {
		score, length := getScore(s[i:])
		result, i = result+score, i+length
	}

	return result
}

func getScore(s string) (score, length int) {
	if s[1] == ')' {
		return 1, 2
	}

	for score, length = 0, 1; s[length] == '('; {
		subScore, subLength := getScore(s[length:])
		score, length = score+subScore, length+subLength
	}

	return score * 2, length + 1
}
