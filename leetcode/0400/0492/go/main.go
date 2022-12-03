package main

func constructRectangle(area int) []int {
	a, b := getDividers(area)
	if b > a {
		a, b = b, a
	}

	return []int{a, b}
}

func getDividers(x int) (int, int) {
	for i := sqrt(x); ; i-- {
		if x%i == 0 {
			return i, x / i
		}
	}
}

func sqrt(x int) int {
	left, right := 0, 3163
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
