package main

import (
	"sort"
)

func frequencySort(s string) string {
	characters := make([]Character, 123)
	for i := 0; i < len(characters); i++ {
		characters[i].Code = byte(i)
	}

	for i := 0; i < len(s); i++ {
		characters[s[i]].Frequency++
	}

	sort.Slice(characters, func(i, j int) bool {
		if characters[i].Frequency == characters[j].Frequency {
			return characters[i].Code > characters[j].Code
		}

		return characters[i].Frequency > characters[j].Frequency
	})

	result := make([]byte, 0, len(s))
	for i := 0; i < len(characters); i++ {
		for ; characters[i].Frequency > 0; characters[i].Frequency-- {
			result = append(result, characters[i].Code)
		}
	}

	return string(result)
}

type Character struct {
	Code      byte
	Frequency int
}
