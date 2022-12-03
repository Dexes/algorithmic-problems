package main

func numDifferentIntegers(word string) int {
	numbers := make(map[string]bool)
	for start, i := 0, 0; i < len(word); i++ {
		if word[i] < '0' || word[i] > '9' {
			continue
		}

		for start = i; start < len(word) && word[start] == '0'; start++ {
		}

		for i++; i < len(word) && '0' <= word[i] && word[i] <= '9'; i++ {
		}

		numbers[word[start:i]] = true
	}

	return len(numbers)
}
