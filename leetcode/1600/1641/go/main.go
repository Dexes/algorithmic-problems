package main

func countVowelStrings(n int) int {
	sums := []int{5, 4, 3, 2, 1}
	for i := 1; i < n; i++ {
		for j := 4; j > 0; j-- {
			sums[j-1] += sums[j]
		}
	}

	return sums[0]
}
