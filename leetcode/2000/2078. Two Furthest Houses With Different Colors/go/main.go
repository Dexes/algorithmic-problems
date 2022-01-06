package main

func maxDistance(colors []int) int {
	left, right := 0, 0

	for i := 1; i < len(colors); i++ {
		if colors[0] == colors[i] {
			continue
		}

		right = i
		if left == 0 {
			left = i
		}
	}

	return max(right, len(colors)-left-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
