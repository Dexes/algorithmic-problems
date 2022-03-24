package main

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	result, x, r := 0, 0, 0
	for left, right := 0, len(people)-1; left <= right; left, right = left+1, r-1 {
		for x, r = limit-people[left], right; left < r && people[r] > x; r-- {
		}

		result += right - r + 1
	}

	return result
}
