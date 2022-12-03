package main

func largestAltitude(gain []int) int {
	result, current := 0, 0
	for i := 0; i < len(gain); i++ {
		current += gain[i]
		if current > result {
			result = current
		}
	}

	return result
}
