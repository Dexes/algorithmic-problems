package main

func isRobotBounded(instructions string) bool {
	directions, direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}, 0
	x, y := 0, 0

	for i := 0; i < 4; i++ {
		for j := 0; j < len(instructions); j++ {
			switch instructions[j] {
			case 'L':
				direction = direction - 1
				if direction < 0 {
					direction = 3
				}
			case 'R':
				direction = direction + 1
				if direction > 3 {
					direction = 0
				}
			default:
				x += directions[direction][0]
				y += directions[direction][1]
			}
		}

		if x == 0 && y == 0 {
			return true
		}
	}

	return false
}
