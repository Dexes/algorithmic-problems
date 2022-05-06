package main

func removeDuplicates(s string, k int) string {
	result, resultIndex := make([]byte, len(s)), -1
	counter, counterIndex := make([]int, len(s)), -1

	for i, x := 0, k-1; i < len(s); i++ {
		switch {
		case resultIndex < 0 || result[resultIndex] != s[i]:
			resultIndex, counterIndex = resultIndex+1, counterIndex+1
			result[resultIndex], counter[counterIndex] = s[i], 1

		case counter[counterIndex] == x:
			resultIndex, counterIndex = resultIndex-x, counterIndex-1

		default:
			resultIndex++
			result[resultIndex], counter[counterIndex] = s[i], counter[counterIndex]+1
		}
	}

	return string(result[:resultIndex+1])
}
