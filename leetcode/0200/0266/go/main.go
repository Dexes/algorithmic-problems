package main

func canPermutePalindrome(s string) bool {
	counter := make([]int, 123)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	for i, oddCount := 'a', 0; i <= 'z'; i++ {
		if oddCount += counter[i] & 1; oddCount > 1 {
			return false
		}
	}

	return true
}
