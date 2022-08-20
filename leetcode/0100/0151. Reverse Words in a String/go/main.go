package main

import "strings"

func reverseWords(s string) string {
	sb, leftIndex := strings.Builder{}, 0
	for rightIndex := len(s) - 1; rightIndex >= 0; rightIndex = leftIndex {
		for ; rightIndex >= 0 && s[rightIndex] == ' '; rightIndex-- {
		}

		for leftIndex = rightIndex; leftIndex >= 0 && s[leftIndex] != ' '; leftIndex-- {
		}

		if leftIndex != rightIndex {
			sb.WriteString(s[leftIndex+1 : rightIndex+1])
			sb.WriteByte(' ')
		}
	}

	return sb.String()[:sb.Len()-1]
}
