package main

func maxLengthBetweenEqualCharacters(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j > i; j-- {
			if j-i <= result {
				break
			}

			if s[i] == s[j] {
				result = j - i
				break
			}
		}
	}

	return result - 1
}
