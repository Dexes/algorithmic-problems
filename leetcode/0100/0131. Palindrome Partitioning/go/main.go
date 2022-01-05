package main

func partition(s string) [][]string {
	return permutations(0, []byte(s), make([]string, 0, 16), make([][]string, 0, 32768))
}

func permutations(start int, data []byte, current []string, result [][]string) [][]string {
	if start == len(data) {
		return append(result, append(make([]string, 0, len(current)), current...))
	}

	for i := start; i < len(data); i++ {
		if isPalindrome(data, start, i) {
			result = permutations(i+1, data, append(current, string(data[start:i+1])), result)
		}
	}

	return result
}

func isPalindrome(s []byte, left, right int) bool {
	for left < right {
		if s[left] == s[right] {
			left, right = left+1, right-1
			continue
		}

		return false
	}

	return true
}
