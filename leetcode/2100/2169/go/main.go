package main

func countOperations(num1 int, num2 int) int {
	result := 0
	for num2 > 0 && num1 > 0 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}

		result++
	}

	return result
}
