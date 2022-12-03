package main

func countGoodTriplets(arr []int, a int, b int, c int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}

			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					result++
				}
			}
		}
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
