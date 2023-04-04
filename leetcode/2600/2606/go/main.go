package main

func maximumCostSubstring(s, chars string, values []int) int {
	var costs [123]int
	for i := int('a'); i <= 'z'; i++ {
		costs[i] = i - 'a' + 1
	}

	for i := 0; i < len(chars); i++ {
		costs[chars[i]] = values[i]
	}

	current, result := 0, 0
	for i := 0; i < len(s); i++ {
		if current += costs[s[i]]; current < 0 {
			current = 0
			continue
		}

		if current > result {
			result = current
		}
	}

	return result
}
