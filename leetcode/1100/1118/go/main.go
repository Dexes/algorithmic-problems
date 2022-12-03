package main

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func numberOfDays(year int, month int) int {
	result := days[month-1]
	if month == 2 && isLeapYear(year) {
		result++
	}

	return result
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
