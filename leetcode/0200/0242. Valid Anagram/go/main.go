package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
		letters[t[i]]--
	}

	for _, count := range letters {
		if count != 0 {
			return false
		}
	}

	return true
}
