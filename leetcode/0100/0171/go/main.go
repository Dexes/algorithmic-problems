package main

func titleToNumber(columnTitle string) int {
	result, base := 0, 1
	for i := len(columnTitle) - 1; i >= 0; i-- {
		result += int(columnTitle[i]-'A'+1) * base
		base *= 26
	}

	return result
}
