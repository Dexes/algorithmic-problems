package main

import (
	"sort"
)

func minimumSum(num int) int {
	digits := []int{num % 10, num / 10 % 10, num / 100 % 10, num / 1000}
	sort.Ints(digits)

	return (digits[0]+digits[1])*10 + digits[2] + digits[3]
}
