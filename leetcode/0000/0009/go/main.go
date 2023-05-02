package main

func isPalindrome(x int) bool {
	if x < 10 {
		return x >= 0
	}

	var reversed int
	for original := x; original > 0; original /= 10 {
		reversed = reversed*10 + original%10
	}

	return reversed == x
}
