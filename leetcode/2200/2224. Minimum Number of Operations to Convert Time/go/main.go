package main

var numbers = []int{60, 15, 5, 1}

func convertTime(current string, correct string) int {
	diff := toMinutes(correct) - toMinutes(current)
	if diff < 0 {
		diff += 1440
	}

	result := 0
	for i := 0; i < len(numbers) && diff > 0; i++ {
		result += diff / numbers[i]
		diff %= numbers[i]
	}

	return result
}

func toMinutes(s string) int {
	return toNumber(s[0], s[1])*60 + toNumber(s[3], s[4])
}

func toNumber(a, b byte) int {
	return int((a-'0')*10 + (b - '0'))
}
