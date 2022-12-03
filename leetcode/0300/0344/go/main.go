package main

func reverseString(s []byte) {
	length := len(s)
	border := length / 2

	for i := 0; i < border; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}
