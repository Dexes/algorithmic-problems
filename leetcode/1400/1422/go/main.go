package main

func maxScore(s string) int {
	aScore, bScore, i := countScore(s)
	result, j, lastIndex := aScore+bScore, 0, len(s)-1
	for i < lastIndex {
		for j = i; j < lastIndex && s[j] == '1'; j++ {
		}
		bScore, i = bScore-(j-i), j

		for ; j < lastIndex && s[j] == '0'; j++ {
		}
		aScore, i = aScore+(j-i), j

		if aScore+bScore > result {
			result = aScore + bScore
		}
	}

	return result
}

func countScore(s string) (int, int, int) {
	aScore, bScore, lastIndex := 0, 0, len(s)-1
	for ; aScore < lastIndex && s[aScore] == '0'; aScore++ {
	}

	i := aScore
	if i == 0 {
		i = 1
	}

	for j := i; j < len(s); j++ {
		if s[j] == '1' {
			bScore++
		}
	}

	return aScore, bScore, i
}
