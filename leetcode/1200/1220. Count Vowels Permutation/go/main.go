package main

const modulo = 1000000007

func countVowelPermutation(n int) int {
	sums, buffer := [5]int{1, 1, 1, 1, 1}, [5]int{}
	for i := 1; i < n; i++ {
		buffer[0] = sums[1]
		buffer[1] = (sums[0] + sums[2]) % modulo
		buffer[2] = (sums[0] + sums[1] + sums[3] + sums[4]) % modulo
		buffer[3] = (sums[2] + sums[4]) % modulo
		buffer[4] = sums[0]

		sums = buffer
	}

	return (sums[0] + sums[1] + sums[2] + sums[3] + sums[4]) % modulo
}
