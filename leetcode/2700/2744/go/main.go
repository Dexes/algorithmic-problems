package main

const capacity = ((('z' - 'a') << 5) | ('z' - 'a')) + 1

func maximumNumberOfStringPairs(words []string) (result int) {
	var counter [capacity]bool
	for _, w := range words {
		if counter[id(w[1], w[0])] {
			result++
		}

		counter[id(w[0], w[1])] = true
	}

	return result
}

func id(a, b byte) int {
	return (int(a-'a') << 5) | int(b-'a')
}
