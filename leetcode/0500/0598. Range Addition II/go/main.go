package main

func maxCount(m int, n int, ops [][]int) int {
	for i := 0; i < len(ops); i++ {
		if ops[i][0] < m {
			m = ops[i][0]
		}

		if ops[i][1] < n {
			n = ops[i][1]
		}
	}

	return m * n
}
