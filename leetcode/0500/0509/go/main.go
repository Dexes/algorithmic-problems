package main

func fib(n int) int {
	if n == 0 {
		return 0
	}

	prev, current := 0, 1
	for i := 1; i < n; i++ {
		prev, current = current, prev+current
	}

	return current
}
