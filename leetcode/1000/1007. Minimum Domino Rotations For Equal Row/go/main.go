package main

func minDominoRotations(tops []int, bottoms []int) int {
	switch a, b := count(tops, bottoms), count(bottoms, tops); {
	case a == -1 && b == -1:
		return -1
	case a == -1:
		return b
	case b == -1:
		return a
	default:
		return min(a, b)
	}
}

func count(first, second []int) int {
	result, same := 0, 0
	for i := 1; i < len(second); i++ {
		switch {
		case first[0] != first[i] && first[0] != second[i]:
			return -1
		case first[0] != first[i]:
			result++
		case first[i] == second[i]:
			same++
		}
	}

	return min(result, len(first)-result-same)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
