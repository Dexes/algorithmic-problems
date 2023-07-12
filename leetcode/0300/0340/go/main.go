package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}

	if len(s) <= k {
		return len(s)
	}

	result, startIndex, letters := 0, -1, [126]int{}
	for i := 0; i < len(s); i++ {
		if letters[s[i]]++; letters[s[i]] > 1 {
			result = max(result, i-startIndex)
			continue
		}

		if k--; k >= 0 {
			result = max(result, i-startIndex)
			continue
		}

		for {
			startIndex++
			if letters[s[startIndex]]--; letters[s[startIndex]] == 0 {
				break
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
