package main

var rotated = []int{1: 1, 6: 9, 8: 8, 9: 6}

func confusingNumber(n int) bool {
	if n == 0 {
		return true
	}

	result, digit := 0, 0
	for x := n; x > 0; x /= 10 {
		if digit = x % 10; digit > 0 && rotated[digit] == 0 {
			return false
		}

		result = result*10 + rotated[digit]
	}

	return result != n
}
