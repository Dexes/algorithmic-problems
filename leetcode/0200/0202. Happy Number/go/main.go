package main

func isHappy(n int) bool {
	numbers := map[int]bool{n: true}

	for n != 1 {
		n = sum(n)
		if numbers[n] {
			return false
		}

		numbers[n] = true
	}

	return true
}

func sum(n int) int {
	result := 0
	for n > 0 {
		result += (n % 10) * (n % 10)
		n /= 10
	}

	return result
}
