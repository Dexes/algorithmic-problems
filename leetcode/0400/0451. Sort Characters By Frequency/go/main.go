package main

import (
	"sort"
)

type Element struct {
	Symbol    byte
	Frequency int
}

func frequencySort(s string) string {
	elements := count(s)
	sort.Slice(elements, func(i, j int) bool {
		if elements[i].Frequency == elements[j].Frequency {
			return elements[i].Symbol > elements[j].Symbol
		}

		return elements[i].Frequency > elements[j].Frequency
	})

	result := make([]byte, 0, len(s))
	for i := 0; i < len(elements); i++ {
		for ; elements[i].Frequency > 0; elements[i].Frequency-- {
			result = append(result, elements[i].Symbol)
		}
	}

	return string(result)
}

func count(s string) []Element {
	counter := make([]int, 255)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	result := make([]Element, 62)
	for i := byte('0'); i <= '9'; i++ {
		result[i-'0'] = Element{i, counter[i]}
	}

	for i := byte('A'); i <= 'Z'; i++ {
		result[i+10-'A'] = Element{i, counter[i]}
	}

	for i := byte('a'); i <= 'z'; i++ {
		result[i+36-'a'] = Element{i, counter[i]}
	}

	return result
}
