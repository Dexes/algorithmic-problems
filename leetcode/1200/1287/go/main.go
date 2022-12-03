package main

func findSpecialInteger(arr []int) int {
	limit, counter := len(arr)/4+1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1] != arr[i] {
			counter = 1
			continue
		}

		counter++
		if counter == limit {
			return arr[i]
		}
	}

	return arr[0]
}
