package main

import (
	"strconv"
)

func fractionAddition(expression string) string {
	var numerator1, denominator1, numerator2, denominator2, i int
	numerator1, denominator1, i = readFraction(expression, 0)
	for i < len(expression) {
		numerator2, denominator2, i = readFraction(expression, i)
		if denominator1 == denominator2 {
			numerator1 += numerator2
		} else {
			numerator1 = (numerator1 * denominator2) + (numerator2 * denominator1)
			denominator1 = denominator1 * denominator2
		}
	}

	numerator1, denominator1 = simplify(numerator1, denominator1)
	return strconv.Itoa(numerator1) + "/" + strconv.Itoa(denominator1)
}

func readFraction(s string, index int) (int, int, int) {
	sign, numerator, denominator := 1, 0, 0

	if s[index] == '-' {
		sign, index = -1, index+1
	} else if s[index] == '+' {
		index++
	}

	if s[index+1] == '0' {
		numerator, index = sign*10, index+3
	} else {
		numerator, index = sign*int(s[index]-'0'), index+2
	}

	if index+1 < len(s) && s[index+1] == '0' {
		denominator, index = 10, index+2
	} else {
		denominator, index = int(s[index]-'0'), index+1
	}

	return numerator, denominator, index
}

func simplify(numerator, denominator int) (int, int) {
	a, b := abs(numerator), denominator
	for b != 0 {
		a, b = b, a%b
	}

	return numerator / a, denominator / a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
