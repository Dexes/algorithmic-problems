package main

func longestPalindrome(words []string) int {
	counter, palindromesCount, result := [26][26]int{}, 0, 0
	for _, w := range words {
		a, b := w[0]-'a', w[1]-'a'
		if counter[b][a] > 0 {
			counter[b][a]--
			result += 4
			if a == b {
				palindromesCount--
			}

			continue
		}

		counter[a][b]++
		if a == b {
			palindromesCount++
		}
	}

	if palindromesCount > 0 {
		return result + 2
	}

	return result
}
