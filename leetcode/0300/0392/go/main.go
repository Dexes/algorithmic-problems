package main

func isSubsequence(s string, t string) bool {
	sIndex, tIndex := 0, 0
	for ; sIndex < len(s); sIndex++ {
		for ; tIndex < len(t) && t[tIndex] != s[sIndex]; tIndex++ {
		}

		if tIndex == len(t) || t[tIndex] != s[sIndex] {
			return false
		}

		tIndex++
	}

	return true
}
