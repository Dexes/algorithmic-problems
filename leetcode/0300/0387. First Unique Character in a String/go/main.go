package main

func firstUniqChar(s string) int {
	indices := getUniqueIndices(s)

	result := len(s) + 1
	for i := 0; i < len(indices); i++ {
		if indices[i] > 0 && indices[i] < result {
			result = indices[i]
		}
	}

	if result == len(s)+1 {
		return -1
	}

	return result - 1
}

func getUniqueIndices(s string) []int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		letter := s[i] - 'a'
		if result[letter] == 0 {
			result[letter] = i + 1
		} else if result[letter] > 0 {
			result[letter] = -1
		}
	}

	return result
}
