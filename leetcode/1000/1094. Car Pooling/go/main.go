package main

func carPooling(trips [][]int, capacity int) bool {
	counter := make([]int, 1001)
	for i := 0; i < len(trips); i++ {
		counter[trips[i][1]] += trips[i][0]
		counter[trips[i][2]] -= trips[i][0]
	}

	for sum, i := 0, 0; i < len(counter); i++ {
		sum += counter[i]
		if sum > capacity {
			return false
		}
	}

	return true
}
