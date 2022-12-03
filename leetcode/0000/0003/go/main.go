package main

func lengthOfLongestSubstring(str string) int {
	result, left, right := 0, 0, 0
	symbols := make([]bool, 127)

	for ; right < len(str); right++ {
		if !symbols[str[right]] {
			symbols[str[right]] = true
			continue
		}

		if width := right - left; width > result {
			result = width
		}

		for ; str[left] != str[right]; left++ {
			symbols[str[left]] = false
		}

		left++
	}

	if width := right - left; width > result {
		result = width
	}

	return result
}
