package main

func maximum69Number(num int) int {
	return num + 3*multiplier(num)
}

func multiplier(num int) int {
	result, current := 0, 1
	for num > 0 {
		if num%10 == 6 {
			result = current
		}

		num /= 10
		current *= 10
	}

	return result
}
