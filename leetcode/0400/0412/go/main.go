package main

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := 14; i < n; i += 15 {
		result[i] = "FizzBuzz"
	}

	for i := 4; i < n; i += 5 {
		if result[i] == "" {
			result[i] = "Buzz"
		}
	}

	for i := 2; i < n; i += 3 {
		if result[i] == "" {
			result[i] = "Fizz"
		}
	}

	for i := 0; i < n; i++ {
		if result[i] == "" {
			result[i] = strconv.Itoa(i + 1)
		}
	}

	return result
}
