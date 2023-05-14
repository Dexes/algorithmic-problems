package main

func circularGameLosers(n int, k int) []int {
	players := make([]bool, n)
	for step, i := 1, 0; !players[i]; step++ {
		players[i] = true
		i = (i + (step * k)) % n
	}

	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if !players[i] {
			result = append(result, i+1)
		}
	}

	return result
}
