package main

import "strings"

func numOfStrings(patterns []string, word string) int {
	result := 0
	for i := 0; i < len(patterns); i++ {
		if strings.Contains(word, patterns[i]) {
			result++
		}
	}

	return result
}
