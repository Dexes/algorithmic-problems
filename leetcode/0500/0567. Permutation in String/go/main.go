package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	ch1, ch2 := countChars(s1), countChars(s2[:len(s1)])
	right := len(s2) - len(s1)
	for i := 0; i < right; i++ {
		if isEqual(ch1, ch2) {
			return true
		}

		ch2[s2[i]-'a']--
		ch2[s2[i+len(s1)]-'a']++
	}

	return isEqual(ch1, ch2)
}

func isEqual(ch1, ch2 []int) bool {
	for i := 0; i < len(ch1); i++ {
		if ch1[i] != ch2[i] {
			return false
		}
	}

	return true
}

func countChars(s string) []int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}

	return result
}
