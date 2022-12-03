package main

func checkDistances(s string, distances []int) bool {
	for i := 0; i < len(s); i++ {
		distance := distances[s[i]-'a'] + 1
		if j := i + distance; j < len(s) && s[j] == s[i] {
			continue
		}

		if j := i - distance; j >= 0 && s[j] == s[i] {
			continue
		}

		return false
	}

	return true
}
