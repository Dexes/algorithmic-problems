package main

func numEquivDominoPairs(dominoes [][]int) int {
	result, counter := 0, count(dominoes)
	for i := 0; i < len(counter); i++ {
		result += counter[i] * (counter[i] - 1) / 2
	}

	return result
}

func count(dominoes [][]int) []int {
	result := make([]int, 81)
	for i := 0; i < len(dominoes); i++ {
		if dominoes[i][0] > dominoes[i][1] {
			dominoes[i][0], dominoes[i][1] = dominoes[i][1], dominoes[i][0]
		}

		result[(dominoes[i][0]-1)*9+dominoes[i][1]-1]++
	}

	return result
}
