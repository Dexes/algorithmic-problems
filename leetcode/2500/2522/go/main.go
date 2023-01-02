package main

func minimumPartition(s string, k int) int {
	result, current, digit := 1, 0, 0
	for i := 0; i < len(s); i++ {
		digit = int(s[i] - '0')
		if digit > k {
			return -1
		}

		current = current*10 + digit
		if current > k {
			current, result = digit, result+1
		}
	}

	return result
}
