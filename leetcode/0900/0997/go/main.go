package main

func findJudge(n int, trust [][]int) int {
	first, second := make([]int, n+1), make([]int, n+1)
	for i := 0; i < len(trust); i++ {
		first[trust[i][0]]++
		second[trust[i][1]]++
	}

	for i := 1; i <= n; i++ {
		if first[i] == 0 && second[i] == n-1 {
			return i
		}
	}

	return -1
}
