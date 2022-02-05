package main

func hammingDistance(x int, y int) int {
	result := 0
	for ; x > 0 || y > 0; x, y = x>>1, y>>1 {
		if x&1 != y&1 {
			result++
		}
	}

	return result
}
