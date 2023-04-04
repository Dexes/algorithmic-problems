package main

func partitionString(s string) (result int) {
	for i, mask := 0, 0; i < len(s); i++ {
		if bit := 1 << (s[i] - 'a'); (mask & bit) == 0 {
			mask |= bit
		} else {
			result, mask = result+1, bit
		}
	}

	return result + 1
}
