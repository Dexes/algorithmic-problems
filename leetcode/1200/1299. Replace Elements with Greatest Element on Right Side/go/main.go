package main

func replaceElements(arr []int) []int {
	max, current := arr[len(arr)-1], 0
	arr[len(arr)-1] = -1
	for i := len(arr) - 2; i >= 0; i-- {
		current, arr[i] = arr[i], max
		if current > max {
			max = current
		}
	}

	return arr
}
