package main

func lengthOfLongestSubstring(str string) int {
	result := 0
	length := len(str)
	left, right := 0, 0
	set := make(map[uint8]bool)

	for ; right < length; right++ {
		if set[str[right]] {
			result = max(result, right-left)
			left = truncate(str, set, left, str[right])
		}

		set[str[right]] = true
	}

	return max(result, right-left)
}

func truncate(str string, set map[uint8]bool, left int, char uint8) int {
	for ; ; left++ {
		set[str[left]] = false
		if str[left] == char {
			return left + 1
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
