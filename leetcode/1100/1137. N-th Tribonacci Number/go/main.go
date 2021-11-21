package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	first, second, third := 0, 1, 1
	for i := 2; i < n; i++ {
		first, second, third = second, third, first+second+third
	}

	return third
}
