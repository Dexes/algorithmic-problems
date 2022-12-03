package main

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	result, index, isNegative := make([]byte, 10), 9, num < 0
	if isNegative {
		num = -num
	}

	for ; num > 0; index-- {
		result[index] = byte(num%7) + '0'
		num /= 7
	}

	if isNegative {
		result[index] = '-'
		return string(result[index:])
	}

	return string(result[index+1:])
}
