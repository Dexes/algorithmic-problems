package main

func longestBeautifulSubstring(word string) (result int) {
	for l, i := 0, startIndex(word, 0); i < len(word); {
		if l, i = length(word, i); l > result {
			result = l
		}
	}

	return result
}

func length(s string, index int) (int, int) {
	i, count := index+1, 1
	for ; i < len(s) && s[i] >= s[i-1]; i++ {
		if s[i] > s[i-1] {
			count++
		}
	}

	if count < 5 {
		return -1, startIndex(s, i)
	}

	return i - index, startIndex(s, i)
}

func startIndex(s string, index int) int {
	for ; index < len(s) && s[index] != 'a'; index++ {
	}

	return index
}
