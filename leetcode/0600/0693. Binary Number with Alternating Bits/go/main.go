package main

func hasAlternatingBits(n int) bool {
	current, prev := n&1, 0
	n >>= 1

	for n > 0 {
		prev, current, n = current, n&1, n>>1
		if prev == current {
			return false
		}
	}

	return true
}
