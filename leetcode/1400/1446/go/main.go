package main

func maxPower(s string) int {
	result := 0
	for i, j := 0, 0; i < len(s); i = j {
		for j = i + 1; j < len(s) && s[i] == s[j]; j++ {
		}

		if j-i > result {
			result = j - i
		}
	}

	return result
}
