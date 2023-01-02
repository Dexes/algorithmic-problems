package main

type TicTacToe struct {
	n     int
	stats [2]*stat
}

func Constructor(n int) TicTacToe {
	return TicTacToe{n: n, stats: [2]*stat{makeStat(n), makeStat(n)}}
}

func (x TicTacToe) Move(row, col, player int) int {
	s := x.stats[player-1]

	if s.cols[col]++; s.cols[col] == x.n {
		return player
	}

	if s.rows[row]++; s.rows[row] == x.n {
		return player
	}

	if col == row {
		if s.straight++; s.straight == x.n {
			return player
		}
	}

	if x.n-row-1 == col {
		if s.reverse++; s.reverse == x.n {
			return player
		}
	}

	return 0
}

func makeStat(n int) *stat {
	return &stat{rows: make([]int, n), cols: make([]int, n), straight: 0, reverse: 0}
}

type stat struct {
	rows, cols []int
	straight   int
	reverse    int
}
