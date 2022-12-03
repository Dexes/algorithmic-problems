package main

import "strconv"

func toHexspeak(num string) string {
	result, index := make([]byte, len(num)), len(num)-1
	x, _ := strconv.Atoi(num)

	for ; x > 0; x, index = x/16, index-1 {
		switch m := x % 16; {
		case m == 0:
			result[index] = 'O'
		case m == 1:
			result[index] = 'I'
		case m >= 10:
			result[index] = byte(m - 10 + 'A')
		default:
			return "ERROR"
		}
	}

	return string(result[index+1:])
}
