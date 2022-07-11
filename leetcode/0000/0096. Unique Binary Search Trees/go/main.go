package main

func numTrees(n int) int {
	cache := make([]int, n+1)
	cache[0], cache[1] = 1, 1

	for i := 2; i <= n; i++ {
		for left, right := i-1, 0; left >= 0; left, right = left-1, right+1 {
			cache[i] += cache[left] * cache[right]
		}
	}

	return cache[n]
}
