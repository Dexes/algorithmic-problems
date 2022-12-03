package main

var daysBeforeMonth = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	from := max(toDays(arriveAlice), toDays(arriveBob))
	to := min(toDays(leaveAlice), toDays(leaveBob))

	return max(to-from+1, 0)
}

func toDays(date string) int {
	month := (date[0]-'0')*10 + date[1] - '1'
	days := int(date[3]-'0')*10 + int(date[4]-'0')

	return daysBeforeMonth[month] + days
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
