package main

func countVowelSubstrings(word string) int {
	result, start, stop := 0, 0, 0
	for i := 0; i < len(word); i = stop + 1 {
		for start = i; start < len(word) && getVowelId(word[start]) == -1; start++ {
		}

		for stop = start; stop < len(word) && getVowelId(word[stop]) >= 0; stop++ {
		}

		result += countSubstrings(word, start, stop)
	}

	return result
}

func countSubstrings(s string, start, stop int) int {
	if stop-start < 5 {
		return 0
	}

	result := 0
	for i := start; i < stop; i++ {
		mask := uint8(0)
		for j := i; j < stop; j++ {
			if id := getVowelId(s[j]); id >= 0 {
				mask |= 1 << id
				if mask == 0b11111 {
					result++
				}
			} else {
				mask = 0
			}
		}
	}

	return result
}

func getVowelId(b byte) int {
	switch b {
	case 'a':
		return 0
	case 'e':
		return 1
	case 'i':
		return 2
	case 'o':
		return 3
	case 'u':
		return 4
	default:
		return -1
	}
}
