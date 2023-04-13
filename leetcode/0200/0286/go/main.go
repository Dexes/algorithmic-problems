package main

const (
	gate  = 0
	empty = 2147483647
)

var queue [17000][2]int

func wallsAndGates(rooms [][]int) {
	popIndex, pushIndex := initQueue(rooms)

	for distance := 1; popIndex != pushIndex; distance++ {
		popIndex, pushIndex = mark(rooms, distance, popIndex, pushIndex)
	}
}

func mark(rooms [][]int, distance, popIndex, pushIndex int) (int, int) {
	i, j, length := 0, 0, qLen(popIndex, pushIndex)

	for ; length > 0; length-- {
		i, j, popIndex = qPop(popIndex)
		if x := i - 1; x > -1 && rooms[x][j] == empty {
			rooms[x][j] = distance
			pushIndex = qPush(x, j, pushIndex)
		}

		if x := i + 1; x < len(rooms) && rooms[x][j] == empty {
			rooms[x][j] = distance
			pushIndex = qPush(x, j, pushIndex)
		}

		if x := j - 1; x > -1 && rooms[i][x] == empty {
			rooms[i][x] = distance
			pushIndex = qPush(i, x, pushIndex)
		}

		if x := j + 1; x < len(rooms[i]) && rooms[i][x] == empty {
			rooms[i][x] = distance
			pushIndex = qPush(i, x, pushIndex)
		}
	}

	return popIndex, pushIndex
}

func initQueue(rooms [][]int) (popIndex, pushIndex int) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			if rooms[i][j] == gate {
				queue[pushIndex][0], queue[pushIndex][1] = i, j
				pushIndex++
			}
		}
	}

	return popIndex, pushIndex
}

func qPush(i, j, index int) int {
	queue[index][0], queue[index][1] = i, j
	if index++; index == len(queue) {
		return 0
	}

	return index
}

func qPop(index int) (int, int, int) {
	i, j := queue[index][0], queue[index][1]
	if index++; index == len(queue) {
		return i, j, 0
	}

	return i, j, index
}

func qLen(popIndex, pushIndex int) (result int) {
	if result = pushIndex - popIndex; result < 0 {
		result += len(queue)
	}

	return result
}
