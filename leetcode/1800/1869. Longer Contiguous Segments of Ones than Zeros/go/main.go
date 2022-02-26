package main

func checkZeroOnes(s string) bool {
	zeros, ones := 0, 0
	for i, j := 0, 0; i < len(s); i = j {
		for j = i + 1; j < len(s) && s[j] == s[i]; j++ {
		}

		switch {
		case s[i] == '0' && j-i > zeros:
			zeros = j - i
		case s[i] == '1' && j-i > ones:
			ones = j - i
		}
	}

	return ones > zeros
}
