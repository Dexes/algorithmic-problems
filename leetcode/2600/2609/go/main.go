package main

func findTheLongestBalancedSubstring(s string) (result int) {
	var i, j, k int
	for ; i < len(s); i = k {
		for j = i; j < len(s) && s[j] == '0'; j++ {
		}

		for k = j; k < len(s) && s[k] == '1'; k++ {
		}

		if x := min(j-i, k-j); x > result {
			result = x
		}
	}

	return result << 1
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
