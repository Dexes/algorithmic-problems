package main

import "strings"

func stringMatching(words []string) []string {
	result := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}

			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
	}

	return result
}
