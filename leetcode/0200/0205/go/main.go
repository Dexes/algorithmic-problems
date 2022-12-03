package main

func isIsomorphic(s string, t string) bool {
	sMap, tMap := make([]byte, 128), make([]byte, 128)
	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != 0 && sMap[s[i]] != t[i] {
			return false
		}

		if tMap[t[i]] != 0 && tMap[t[i]] != s[i] {
			return false
		}

		sMap[s[i]], tMap[t[i]] = t[i], s[i]
	}

	return true
}
