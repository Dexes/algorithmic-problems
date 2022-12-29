package main

func reverseWords(s []byte) {
	reverse(s, 0, len(s)-1)

	for i, j := 0, spaceIndex(s, 1); i < len(s); i, j = j+1, spaceIndex(s, j+2) {
		reverse(s, i, j-1)
	}
}

func spaceIndex(s []byte, index int) int {
	for ; index < len(s) && s[index] != ' '; index++ {
	}

	return index
}

func reverse(s []byte, left, right int) {
	for ; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}
