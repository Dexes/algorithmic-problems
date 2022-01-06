package main

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	result := 0
	for i := 0; i < len(seats); i++ {
		result += abs(seats[i] - students[i])
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
