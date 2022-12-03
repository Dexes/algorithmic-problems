package main

func arrangeCoins(n int) int {
	left, right := 0, 100000
	for {
		middle := left + (right-left)/2
		count := (middle + 1) * middle / 2
		if count <= n && n <= count+middle {
			return middle
		}

		if count > n {
			right = middle
		} else {
			left = middle
		}
	}
}
