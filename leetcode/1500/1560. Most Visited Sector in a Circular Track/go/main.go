package main

func mostVisited(n int, rounds []int) []int {
	first, last := rounds[0], rounds[len(rounds)-1]
	if first <= last {
		result := make([]int, 0, last-first+1)
		return fill(result, first, last)
	}

	result := make([]int, 0, n-first+last+1)
	result = fill(result, 1, last)
	return fill(result, first, n)
}

func fill(arr []int, first, last int) []int {
	for ; first <= last; first++ {
		arr = append(arr, first)
	}

	return arr
}
