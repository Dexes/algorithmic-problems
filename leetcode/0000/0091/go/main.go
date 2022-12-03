package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	if isInvalid(s[0], s[1]) {
		return 0
	}

	dp := make([]int, len(s))
	if isAmbiguous(s[0], s[1]) {
		dp[0], dp[1] = 1, 2
	} else {
		dp[0], dp[1] = 1, 1
	}

	for i := 2; i < len(s); i++ {
		if isInvalid(s[i-1], s[i]) {
			return 0
		}

		if s[i] == '0' {
			dp[i] = dp[i-2]
			continue
		}

		if isAmbiguous(s[i-1], s[i]) {
			dp[i] = dp[i-1] + dp[i-2]
			continue
		}

		dp[i] = dp[i-1]
	}

	return dp[len(s)-1]
}

func isInvalid(prev, current byte) bool {
	return current == '0' && (prev < '1' || prev > '2')
}

func isAmbiguous(prev, current byte) bool {
	if current == '0' {
		return false
	}

	return prev == '1' || (prev == '2' && current < '7')
}
