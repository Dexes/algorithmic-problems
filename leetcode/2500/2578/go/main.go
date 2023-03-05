package main

import "sort"

func splitNum(num int) int {
	digits := make([]int, 0, 10)
	for ; num > 0; num /= 10 {
		digits = append(digits, num%10)
	}
	sort.Ints(digits)

	a, b := 0, 0
	for i := 0; i < len(digits); i++ {
		a, b = b*10+digits[i], a
	}

	return a + b
}
