package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	result := 0
	for i, j := 0, -1; i < len(g); i++ {
		for j++; j < len(s) && s[j] < g[i]; j++ {
		}

		if j == len(s) {
			return result
		}

		result++
	}

	return result
}
