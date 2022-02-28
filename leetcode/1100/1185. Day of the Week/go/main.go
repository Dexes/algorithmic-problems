package main

var days = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var names = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func dayOfTheWeek(day int, month int, year int) string {
	return names[(countDays(day, month, year)+4)%7]
}

func countDays(day int, month int, year int) int {
	result := 0
	for i := 1971; i < year; i++ {
		if isLeapYear(i) {
			result += 366
		} else {
			result += 365
		}
	}

	return result + dayOfYear(day, month, year)
}

func dayOfYear(day int, month int, year int) int {
	for i := 1; i < month; i++ {
		day += days[i-1]
	}

	if isLeapYear(year) && month > 2 {
		day++
	}

	return day
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
