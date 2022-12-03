package main

func findLucky(arr []int) int {
	counter := make([]int, 501)
	for i := 0; i < len(arr); i++ {
		counter[arr[i]]++
	}

	for i := len(counter) - 1; i > 0; i-- {
		if counter[i] == i {
			return i
		}
	}

	return -1
}
