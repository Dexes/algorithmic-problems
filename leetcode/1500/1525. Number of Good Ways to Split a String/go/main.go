package main

func numSplits(s string) (result int) {
	lLetters, lLen, rLetters, rLen := make([]int, 123), 0, make([]int, 123), 0
	for i := 0; i < len(s); i++ {
		if rLetters[s[i]]++; rLetters[s[i]] == 1 {
			rLen++
		}
	}

	for i := 0; i < len(s); i++ {
		if lLetters[s[i]]++; lLetters[s[i]] == 1 {
			lLen++
		}

		if rLetters[s[i]]--; rLetters[s[i]] == 0 {
			rLen--
		}

		if lLen == rLen {
			result++
		}
	}

	return result
}
