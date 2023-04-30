package main

func isWinner(player1, player2 []int) int {
	score1, score2, double1, double2 := 0, 0, 0, 0
	for i := 0; i < len(player1); i++ {
		score1, double1 = increaseScore(player1[i], score1, double1)
		score2, double2 = increaseScore(player2[i], score2, double2)
	}

	if score1 > score2 {
		return 1
	}

	if score2 > score1 {
		return 2
	}

	return 0
}

func increaseScore(pins, score, double int) (int, int) {
	if double == 0 {
		score += pins
	} else {
		score += pins * 2
		double--
	}

	if pins == 10 {
		return score, 2
	}

	return score, double
}
