package main

func minDistance(word1 string, word2 string) int {
	return len(word1) + len(word2) - lcs(word1, word2)<<1
}

func lcs(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				continue
			}

			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[len(s1)][len(s2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
