package main

func sequentialDigits(low int, high int) []int {
	result, length := make([]int, 0, 36), countDigits(low)
	start, diff := getSequenceStart(length), getSequenceDiff(length)
	for start < low && start%10 > 0 {
		start += diff
	}

	for {
		for x := start; x%10 > 0; x += diff {
			if x > high {
				return result
			}

			result = append(result, x)
		}

		length++
		start, diff = getSequenceStart(length), diff*10+1
	}
}

func countDigits(n int) int {
	result := 0
	for n > 0 {
		result++
		n /= 10
	}

	return result
}

func getSequenceStart(length int) int {
	result := 1
	for i := 2; i <= length; i++ {
		result = result*10 + i
	}

	return result
}

func getSequenceDiff(length int) int {
	result := 1
	for i := 2; i <= length; i++ {
		result = result*10 + 1
	}

	return result
}
