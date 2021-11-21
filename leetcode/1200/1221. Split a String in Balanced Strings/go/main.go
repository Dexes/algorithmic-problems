package main

func balancedStringSplit(s string) int {
	balance, result := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			balance++
		} else {
			balance--
		}

		if balance == 0 {
			result++
		}
	}

	return result
}
