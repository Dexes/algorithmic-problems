package main

func hasGroupsSizeX(deck []int) bool {
	gcd, numbers := 0, count(deck)
	for _, number := range numbers {
		if gcd = GCD(number, gcd); gcd == 1 {
			return false
		}
	}

	return true
}

func count(nums []int) map[int]int {
	numbers := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numbers[nums[i]]++
	}

	return numbers
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
