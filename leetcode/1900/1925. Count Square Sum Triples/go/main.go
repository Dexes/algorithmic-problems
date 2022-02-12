package main

func countTriples(n int) int {
	result, nSqr := 0, n*n
	for i := 2; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := (i * i) + (j * j)
			if x > nSqr {
				break
			}

			if isSqr(x) {
				result += 2
			}
		}
	}

	return result
}

func isSqr(x int) bool {
	left, right := 0, 62500
	for {
		middle := left + (right-left)/2
		begin, end := middle*middle, (middle+1)*(middle+1)

		if begin <= x && x < end {
			return begin == x
		}

		if x < begin {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
}
