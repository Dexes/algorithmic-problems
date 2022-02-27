package main

func diStringMatch(s string) []int {
	result, low, high := make([]int, len(s)+1), 0, len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			result[i], high = high, high-1
		} else {
			result[i], low = low, low+1
		}
	}

	result[len(s)] = low
	return result
}
