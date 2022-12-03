package main

func repeatedCharacter(s string) byte {
	letters := [26]bool{}
	for i := 0; i < len(s); i++ {
		if letters[s[i]-'a'] {
			return s[i]
		}

		letters[s[i]-'a'] = true
	}

	return 0
}
