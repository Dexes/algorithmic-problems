package main

func isPerfectSquare(num int) bool {
	left, right := 0, num

	for left < right {
		middle := left + (right-left)/2
		square := middle * middle

		if square > num {
			right = middle
			continue
		}

		if square < num {
			left = middle + 1
			continue
		}

		return true
	}

	return num == 1
}
