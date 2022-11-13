package main

import "strings"

func reverseWords(s string) string {
	builder := strings.Builder{}
	for left, right := 0, len(s)-1; right >= 0; right = left {
		for ; right >= 0 && s[right] == ' '; right-- {
		}

		for left = right; left >= 0 && s[left] != ' '; left-- {
		}

		if left != right {
			builder.WriteString(s[left+1 : right+1])
			builder.WriteByte(' ')
		}
	}

	return builder.String()[:builder.Len()-1]
}
