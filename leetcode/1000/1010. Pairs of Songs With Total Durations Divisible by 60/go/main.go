package main

func numPairsDivisibleBy60(time []int) int {
	result, counter := 0, make([]int, 60)
	for i := 0; i < len(time); i++ {
		remainder := time[i] % 60
		result += counter[(60-remainder)%60]
		counter[remainder]++
	}

	return result
}
