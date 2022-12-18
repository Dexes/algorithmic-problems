package main

type SnakeGame struct {
	width, height int
	score         int
	food          [][]int
	body          [][]int
}

func Constructor(width int, height int, food [][]int) SnakeGame {
	body := make([][]int, 1, 51)
	body[0] = []int{0, 0}

	return SnakeGame{width: width, height: height, score: 0, food: food, body: body}
}

func (x *SnakeGame) Move(direction string) int {
	headIndex := len(x.body) - 1
	r, c := nextHead(x.body[headIndex], direction)

	if len(x.food) > 0 && r == x.food[0][0] && c == x.food[0][1] {
		x.body = append(x.body, x.food[0])
		x.food = x.food[1:]
		x.score++
		return x.score
	}

	if r < 0 || r == x.height || c < 0 || c == x.width {
		return -1
	}

	for i := 1; i <= headIndex; i++ {
		if x.body[i][0] == r && x.body[i][1] == c {
			return -1
		}

		x.body[i-1][0], x.body[i-1][1] = x.body[i][0], x.body[i][1]
	}

	x.body[headIndex][0], x.body[headIndex][1] = r, c
	return x.score
}

func nextHead(head []int, direction string) (int, int) {
	if direction[0] == 'U' {
		return head[0] - 1, head[1]
	}

	if direction[0] == 'D' {
		return head[0] + 1, head[1]
	}

	if direction[0] == 'L' {
		return head[0], head[1] - 1
	}

	return head[0], head[1] + 1
}
