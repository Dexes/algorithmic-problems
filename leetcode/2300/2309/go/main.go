package main

func greatestLetter(s string) string {
	lower, upper := [26]bool{}, [26]bool{}
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' {
			upper[s[i]-'A'] = true
		} else {
			lower[s[i]-'a'] = true
		}
	}

	for i := 25; i >= 0; i-- {
		if lower[i] && upper[i] {
			return string(byte(i) + 'A')
		}
	}

	return ""
}
