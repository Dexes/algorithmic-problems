package main

func longestPalindrome(s string) int {
	result, letters, counter := 0, make([]int, 52), 0
	for i := 0; i < len(s); i++ {
		index := letterIndex(s[i])
		letters[index]++
		if letters[index]&1 == 1 {
			counter++
		} else {
			counter--
			result += 2
		}
	}

	if counter > 0 {
		return result + 1
	}

	return result
}

func letterIndex(b byte) byte {
	if b > 'Z' {
		return b - 'a' + 26
	}

	return b - 'A'
}
