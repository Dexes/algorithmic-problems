package main

func mySqrt(x int) int {
	left, right := 0, 46341
	for left < right {
		middle := left + (right-left)/2
		begin, end := middle*middle, (middle+1)*(middle+1)

		if begin <= x && x < end {
			return middle
		}

		if x == end {
			return middle + 1
		}

		if x < begin {
			right = middle
		} else {
			left = middle + 1
		}
	}

	return -1
}
