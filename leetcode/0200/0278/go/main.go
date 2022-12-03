package main

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	left, right := 1, n

	for left < right {
		middle := left + (right-left)/2

		if isBadVersion(middle) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return left
}
