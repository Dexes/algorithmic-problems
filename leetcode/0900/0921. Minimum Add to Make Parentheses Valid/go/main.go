package main

func minAddToMakeValid(s string) int {
	result, balance := 0, 0
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '(':
			balance++
		case balance == 0:
			result++
		default:
			balance--
		}
	}

	return result + balance
}
