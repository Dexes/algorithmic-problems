package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	border := length / 2

	for i := 0; i < border; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}
