package main

func pivotInteger(n int) int {
	left, right, sum := 1, n, n*(n+1)/2
	for left < right {
		if middle := left + (right-left)/2; middle*middle < sum {
			left = middle + 1
		} else {
			right = middle
		}
	}

	if left*left == sum {
		return left
	}

	return -1
}
