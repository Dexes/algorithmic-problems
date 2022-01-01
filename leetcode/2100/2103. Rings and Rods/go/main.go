package main

func countPoints(rings string) int {
	result, masks := 0, make([]int, 10)
	for i := 0; i < len(rings); i += 2 {
		index := rings[i+1] - '0'
		if masks[index] == 0b111 {
			continue
		}

		switch rings[i] {
		case 'R':
			masks[index] |= 0b001
		case 'G':
			masks[index] |= 0b010
		case 'B':
			masks[index] |= 0b100
		}

		if masks[index] == 0b111 {
			result++
		}
	}

	return result
}
