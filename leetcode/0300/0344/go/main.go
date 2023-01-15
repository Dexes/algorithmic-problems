package main

func reverseString(s []byte) {
	rightIndex := len(s) - 1
	for i := (len(s) / 2) - 1; i >= 0; i-- {
		s[i], s[rightIndex-i] = s[rightIndex-i], s[i]
	}
}
