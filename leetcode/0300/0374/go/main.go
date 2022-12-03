package main

func guess(num int) int {
	return 0
}

func guessNumber(right int) int {
	for left := 1; ; {
		middle := (right + left) / 2
		answer := guess(middle)
		if answer == -1 {
			right = middle
			continue
		}

		if answer == 1 {
			left = middle + 1
			continue
		}

		return middle
	}
}
