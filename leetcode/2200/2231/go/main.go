package main

import (
	"sort"
)

func largestInteger(num int) (result int) {
	odd, even, parity := digitsInfo(num)
	oddIndex, evenIndex := len(odd)-1, len(even)-1

	sort.Ints(odd)
	sort.Ints(even)

	for i := len(parity) - 1; i >= 0; i-- {
		result *= 10

		if parity[i] {
			result += odd[oddIndex]
			oddIndex--
		} else {
			result += even[evenIndex]
			evenIndex--
		}
	}

	return result
}

func digitsInfo(num int) ([]int, []int, []bool) {
	odd, even, parity := make([]int, 0, 10), make([]int, 0, 10), make([]bool, 0, 10)
	for ; num > 0; num /= 10 {
		digit := num % 10
		p := digit&1 == 1

		parity = append(parity, p)
		if p {
			odd = append(odd, digit)
		} else {
			even = append(even, digit)
		}
	}

	return odd, even, parity
}
