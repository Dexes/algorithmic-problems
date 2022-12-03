package main

func countSegments(s string) int {
	for result, i := 0, 0; ; {
		for ; i < len(s) && s[i] == ' '; i++ {
		}

		if i == len(s) {
			return result
		}

		for ; i < len(s) && s[i] != ' '; i++ {
		}

		result++
	}
}
