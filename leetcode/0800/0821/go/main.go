package main

func shortestToChar(s string, c byte) []int {
	result, i := make([]int, len(s)), find(s, c, 0)
	for x := 0; x < i; x++ {
		result[x] = i - x
	}

	for j := find(s, c, i+1); j < len(s); i, j = j, find(s, c, j+1) {
		for counter := (j - i) >> 1; counter > 0; counter-- {
			result[i+counter] = counter
			result[j-counter] = counter
		}
	}

	for x := i + 1; x < len(s); x++ {
		result[x] = x - i
	}

	return result
}

func find(s string, c byte, index int) int {
	for ; index < len(s) && s[index] != c; index++ {
	}

	return index
}
