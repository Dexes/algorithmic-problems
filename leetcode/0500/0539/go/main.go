package main

const max = 1440

func findMinDifference(timePoints []string) int {
	if len(timePoints) == 2 {
		left, right := toMinutes(timePoints[0]), toMinutes(timePoints[1])
		return min((right-left+max)%max, (left-right+max)%max)
	}

	var (
		points [max]int

		left, prev, current, result = max, max, max, max
	)

	for i := 0; i < len(timePoints); i++ {
		minutes := toMinutes(timePoints[i])
		if points[minutes]++; points[minutes] == 2 {
			return 0
		}

		if minutes < left {
			left = minutes
		}
	}

	for prev, current = left, left; ; prev = current {
		for current++; current < len(points) && points[current] == 0; current++ {
		}

		if current == len(points) {
			break
		}

		distance := min(current-prev, prev-current+max)
		if distance == 1 {
			return 1
		}

		if distance < result {
			result = distance
		}
	}

	return min(result, min(prev-left, left-prev+max))
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func toMinutes(s string) int {
	return toNumber(s[0], s[1])*60 + toNumber(s[3], s[4])
}

func toNumber(a, b byte) int {
	return int((a-'0')*10 + (b - '0'))
}
