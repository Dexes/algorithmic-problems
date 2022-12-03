package main

func removeAnagrams(words []string) []string {
	rIndex, wIndex := 1, 0
	for ; rIndex < len(words); rIndex++ {
		if !areAnagrams(words[rIndex], words[wIndex]) {
			wIndex++
			words[wIndex] = words[rIndex]
		}
	}

	return words[:wIndex+1]
}

func areAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	aCounter, bCounter := [26]byte{}, [26]byte{}
	for i := range a {
		aCounter[a[i]-'a']++
		bCounter[b[i]-'a']++
	}

	return aCounter == bCounter
}
