package main

func appendCharacters(s string, t string) int {
	for i, j := -1, 0; j < len(t); j++ {
		if i = find(s, t[j], i+1); i == len(s) {
			return len(t) - j
		}
	}

	return 0
}

func find(s string, b byte, i int) int {
	for ; i < len(s) && s[i] != b; i++ {
	}

	return i
}
