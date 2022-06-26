package main

func countAsterisks(s string) int {
	result, toggle := 0, true
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			toggle = !toggle
			continue
		}

		if toggle && s[i] == '*' {
			result++
		}
	}

	return result
}
