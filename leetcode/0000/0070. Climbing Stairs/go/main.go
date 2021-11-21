package main

func climbStairs(n int) int {
	prev, current := 0, 1
	for i := 0; i < n; i++ {
		prev, current = current, prev+current
	}

	return current
}
