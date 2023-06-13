package main

type array [200]int

var nums array

func equalPairs(grid [][]int) (result int) {
	m := make(map[array]int)
	for row := 0; row < len(grid); row++ {
		copy(nums[:], grid[row])
		m[nums]++
	}

	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid); row++ {
			nums[row] = grid[row][col]
		}

		result += m[nums]
	}

	return result
}
