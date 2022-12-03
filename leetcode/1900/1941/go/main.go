package main

func areOccurrencesEqual(s string) bool {
	counter, index := make([]int, 26), 0
	for i := 0; i < len(s); i++ {
		index = int(s[i] - 'a')
		counter[index]++
	}

	for i := 0; i < len(counter); i++ {
		if counter[i] != 0 && counter[i] != counter[index] {
			return false
		}
	}

	return true
}
