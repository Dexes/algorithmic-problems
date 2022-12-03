package main

func areSentencesSimilar(s1 []string, s2 []string, similar [][]string) bool {
	if len(s1) != len(s2) {
		return false
	}

	pairs := make(map[string]bool)
	for _, x := range similar {
		pairs[concat(x[0], x[1])] = true
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] && !pairs[concat(s1[i], s2[i])] && !pairs[concat(s2[i], s1[i])] {
			return false
		}
	}

	return true
}

func concat(s1, s2 string) string {
	return s1 + "|" + s2
}
