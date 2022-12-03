package main

func executeInstructions(n int, startPos []int, s string) []int {
	result := make([]int, len(s))
	for startCommand := 0; startCommand < len(s); startCommand++ {
		row, col := startPos[0], startPos[1]
		counter := 0
		for command := startCommand; command < len(s); command++ {
			row, col = executeCommand(s[command], row, col)
			if row < 0 || row >= n || col < 0 || col >= n {
				break
			}

			counter++
		}

		result[startCommand] = counter
	}

	return result
}

func executeCommand(command byte, row, col int) (int, int) {
	switch command {
	case 'L':
		col--
	case 'R':
		col++
	case 'U':
		row--
	case 'D':
		row++
	}

	return row, col
}
