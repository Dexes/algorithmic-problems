package main

func mySqrt(x int) int {
	left, right := 0, 46341
	for {
		middle := left + (right-left)/2
		begin, end := middle*middle, (middle+1)*(middle+1)

		if begin <= x && x < end {
			return middle
		}

		if x < begin {
			right = middle
		} else {
			left = middle + 1
		}
	}
}
