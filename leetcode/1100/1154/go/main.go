package main

func dayOfYear(date string) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	year, month, day := parseDate(date)

	for i := 1; i < month; i++ {
		day += days[i-1]
	}

	if isLeapYear(year) && month > 2 {
		day++
	}

	return day
}

func parseDate(date string) (int, int, int) {
	return (int(date[0])-'0')*1000 + (int(date[1])-'0')*100 + (int(date[2])-'0')*10 + int(date[3]) - '0',
		(int(date[5])-'0')*10 + int(date[6]) - '0',
		(int(date[8])-'0')*10 + int(date[9]) - '0'
}

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
