package main

func getLucky(s string, k int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		number := int(s[i] - 'a' + 1)
		result += number/10 + number%10
	}

	for i := 1; i < k; i++ {
		result = sum(result)
	}

	return result
}

func sum(n int) int {
	result := 0
	for n > 0 {
		result += n % 10
		n /= 10
	}

	return result
}
