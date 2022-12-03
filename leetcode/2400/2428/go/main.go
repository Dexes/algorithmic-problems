package main

func maxSum(grid [][]int) (result int) {
	var r1, r2, r3, cl1, cl2, cr1, cr2 int

	for r3 = 2; r3 < len(grid); r3++ {
		r1, r2 = r3-2, r3-1

		sum := grid[r1][0] + grid[r1][1] + grid[r1][2] + grid[r2][1] + grid[r3][0] + grid[r3][1] + grid[r3][2]
		result = max(sum, result)
		for cr2 = 3; cr2 < len(grid[r1]); cr2++ {
			cl1, cl2, cr1 = cr2-3, cr2-2, cr2-1

			sum += grid[r1][cr2] + grid[r2][cr1] + grid[r3][cr2] - grid[r1][cl1] - grid[r2][cl2] - grid[r3][cl1]
			result = max(sum, result)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
