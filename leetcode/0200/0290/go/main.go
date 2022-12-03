package main

func wordPattern(pattern string, s string) bool {
	pMap, sMap := make([]string, 26), make(map[string]byte)
	wordStart := 0
	for i := 0; i < len(pattern); i++ {
		wordEnd := findSpace(s, wordStart)
		if wordStart == wordEnd {
			return false
		}

		pIndex, sIndex := pattern[i]+1-'a', s[wordStart:wordEnd]
		if pMap[pIndex] != "" && pMap[pIndex] != sIndex {
			return false
		}

		if sMap[sIndex] != 0 && sMap[sIndex] != pIndex {
			return false
		}

		pMap[pIndex], sMap[sIndex] = sIndex, pIndex
		wordStart = wordEnd + 1
	}

	return wordStart == len(s)+1
}

func findSpace(s string, index int) int {
	for ; index < len(s) && s[index] != ' '; index++ {
	}

	return index
}
