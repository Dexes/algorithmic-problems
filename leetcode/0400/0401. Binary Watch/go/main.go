package main

func readBinaryWatch(turnedOn int) []string {
	return makePermutations(0, 0, turnedOn, make([]string, 0, 200))
}

func makePermutations(mask uint16, index, turnedOn int, permutations []string) []string {
	if turnedOn == 0 {
		if time, err := toTime(mask); !err {
			return append(permutations, time)
		}

		return permutations
	}

	if index+turnedOn > 10 {
		return permutations
	}

	permutations = makePermutations(mask|1<<index, index+1, turnedOn-1, permutations)
	permutations = makePermutations(mask, index+1, turnedOn, permutations)

	return permutations
}

func toTime(mask uint16) (string, bool) {
	hours, minutes := byte(0), byte(0)
	hoursMask, minutesMask := byte(mask>>6), byte(mask)
	for i := 0; i < 6; i++ {
		bit := byte(1) << i
		if hoursMask&bit == bit {
			hours += bit
		}

		if minutesMask&bit == bit {
			minutes += bit
		}
	}

	if hours > 11 || minutes > 59 {
		return "", true
	}

	return toString(hours, minutes), false
}

func toString(hours, minutes byte) string {
	result := make([]byte, 0, 5)

	if hours == 10 || hours == 11 {
		result = append(result, '1')
	}
	result = append(result, '0'+hours%10, ':')

	if minutes < 10 {
		result = append(result, '0', '0'+minutes)
	} else {
		result = append(result, '0'+minutes/10, '0'+minutes%10)
	}

	return string(result)
}
