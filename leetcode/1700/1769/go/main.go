package main

func minOperations(boxes string) []int {
	prefix := prefixSub(boxes)

	result := make([]int, len(boxes))
	for i := 1; i < len(boxes); i++ {
		result[i] = result[i-1] + prefix[i-1]
	}

	balls := prefix[len(prefix)-1]
	for i := len(boxes) - 2; i >= 0; i-- {
		result[i] += result[i+1] - (result[i] + prefix[i]) + (balls - prefix[i])
	}

	return result
}

func prefixSub(s string) []int {
	result := make([]int, len(s))
	if s[0] == '1' {
		result[0] = 1
	}

	for i := 1; i < len(s); i++ {
		result[i] = result[i-1]
		if s[i] == '1' {
			result[i]++
		}
	}

	return result
}
