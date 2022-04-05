package main

func minBitFlips(start int, goal int) (result int) {
	for ; start > 0 || goal > 0; start, goal = start>>1, goal>>1 {
		if start&1 != goal&1 {
			result++
		}
	}

	return result
}
