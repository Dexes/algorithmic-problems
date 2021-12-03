package main

func maxDepth(s string) int {
	result, balance := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			balance++
			if balance > result {
				result = balance
			}
		} else if s[i] == ')' {
			balance--
		}
	}

	return result
}
