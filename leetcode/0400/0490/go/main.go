package main

const (
	left = 1 << (iota + 1)
	right
	up
	down
)

func hasPath(maze [][]int, start, destination []int) bool {
	return run(left, maze, start[0], start[1]-1, destination) ||
		run(right, maze, start[0], start[1]+1, destination) ||
		run(up, maze, start[0]-1, start[1], destination) ||
		run(down, maze, start[0]+1, start[1], destination)
}

func run(direction int, maze [][]int, row, col int, destination []int) bool {
	if isWall(maze, row, col) {
		return false
	}

	if (maze[row][col] & direction) > 0 {
		return false
	}

	maze[row][col] |= direction

	nextRow, nextCol := next(direction, row, col)
	if isWall(maze, nextRow, nextCol) {
		if row == destination[0] && col == destination[1] {
			return true
		}

		if direction == left || direction == right {
			return run(up, maze, row-1, col, destination) || run(down, maze, row+1, col, destination)
		}

		return run(left, maze, row, col-1, destination) || run(right, maze, row, col+1, destination)
	}

	return run(direction, maze, nextRow, nextCol, destination)
}

func isWall(maze [][]int, row, col int) bool {
	return row < 0 || row == len(maze) || col < 0 || col == len(maze[row]) || maze[row][col] == 1
}

func next(direction, row, col int) (int, int) {
	if direction == left {
		return row, col - 1
	}

	if direction == right {
		return row, col + 1
	}

	if direction == up {
		return row - 1, col
	}

	return row + 1, col
}
