package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letters := make([]int, 133)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
		letters[t[i]]--
	}

	for i := 'a'; i <= 'z'; i++ {
		if letters[i] != 0 {
			return false
		}
	}

	return true
}
